package integratedprovider

import (
	"bytes"
	"context"
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stateprism/libprisma/cryptoutil"
	"github.com/stateprism/libprisma/cryptoutil/encryption"
	"github.com/stateprism/libprisma/cryptoutil/kdf"
	"github.com/stateprism/libprisma/memkv"
	"github.com/stateprism/shell_vault/server/middleware/auth"
	"github.com/stateprism/shell_vault/server/plugins"
	"github.com/stateprism/shell_vault/server/providers"
	"github.com/vmihailenco/msgpack/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"path"
	"time"
)

type LocalProvider struct {
	conf         providers.ConfigurationProvider
	ephemeralKey []byte
	logger       *zap.Logger
	db           *sql.DB
	plugins      *plugins.Provider
}

func New(plugins *plugins.Provider, config providers.ConfigurationProvider, log *zap.Logger) (providers.AuthProvider, error) {
	key := kdf.PbKdf2.Key(cryptoutil.NewRandom(128), 4096, 32, sha512.New)
	localStore, err := config.GetString("paths.data")
	if err != nil {
		return nil, err
	}

	dbPath := path.Join(localStore, "users.db")
	var isDBInit bool
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		isDBInit = true
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if _, err := db.Exec("SELECT principal FROM users where principal = 'root';"); err != nil {
		// Table does not exist
		isDBInit = true
	}

	provider := &LocalProvider{
		conf:         config,
		db:           db,
		ephemeralKey: key.GetKey(),
		logger:       log,
		plugins:      plugins,
	}

	if isDBInit {
		err := provider.InitNewDB()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize database: %s", err)
		}
	}
	return provider, nil
}

func (p *LocalProvider) String() string {
	return "LOCAL"
}

func (p *LocalProvider) Authenticate(ctx context.Context) (string, error) {
	sesTtlCtx := ctx.Value("sessionTtl")
	sesTtl, ok := sesTtlCtx.(time.Duration)
	if sesTtlCtx == nil || !ok {
		panic("invalid session ttl on authentication attempt, check configurations")
	}
	token, err := auth.Metadata(ctx, "local", "authorization")
	if err != nil {
		return "", err
	}

	tokenBytes := make([]byte, len(token))
	_, err = base64.StdEncoding.Decode(tokenBytes, []byte(token))
	if err != nil {
		return "", status.Error(codes.Unauthenticated, "Your request is not acceptable by any of the enabled auth providers")
	}
	username, password, found := bytes.Cut(tokenBytes, []byte{0x1E})
	if !found {
		return "", status.Error(codes.Unauthenticated, "Your request is not acceptable by any of the enabled auth providers")
	}

	username = bytes.ReplaceAll(username, []byte{0x00}, []byte{})
	password = bytes.ReplaceAll(password, []byte{0x00}, []byte{})
	// Look up user in the database
	var data []byte
	err = p.db.QueryRow(GetExtraDataQuery, string(username)).Scan(&data)
	if err != nil {
		time.Sleep(1 * time.Second)
		return "", status.Error(codes.Unauthenticated, "Invalid username or password")
	}
	// Load the user data
	tbuf := make(map[string]any)
	err = json.Unmarshal(data, &tbuf)
	if err != nil {
		return "", status.Error(codes.Internal, "An internal error occurred")
	}
	entInfo := memkv.NewMemKV(".", nil)
	if err := entInfo.LoadFromSerializableMap(tbuf); err != nil {
		return "", status.Error(codes.Internal, "An internal error occurred")
	}
	// Check the password
	pwString, _ := entInfo.Get("auth.auth_token")
	pwKey, _ := kdf.PbKdf2.FromString(pwString.(string))
	if !pwKey.Equals(string(password)) {
		time.Sleep(1 * time.Second)
		return "", status.Error(codes.Unauthenticated, "Invalid username or password")
	}

	sesId := uuid.New()
	session := &providers.SessionInfo{
		Principal: string(username),
		Realm:     "local",
		Id:        sesId,
		Deadline:  time.Now().Add(sesTtl).UTC().Unix(),
	}
	data, err = msgpack.Marshal(session)
	if err != nil {
		return "", status.Error(codes.Internal, "An internal error occurred")
	}
	encrypted, err := p.encryptSession(data)
	if err != nil {
		return "", status.Error(codes.Internal, "An internal error occurred")
	}
	token = base64.StdEncoding.EncodeToString(encrypted)
	return token, nil
}

func (p *LocalProvider) Authorize(ctx context.Context, method string) (context.Context, error) {
	token, err := auth.Metadata(ctx, "encrypted", "authorization")
	if err != nil {
		return nil, err
	}

	tokenBytes := make([]byte, base64.StdEncoding.DecodedLen(len(token)))
	// trim any null bytes from the right
	n, err := base64.StdEncoding.Decode(tokenBytes, []byte(token))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Your request is not acceptable by any of the enabled auth providers")
	}
	if n < 32 {
		return nil, status.Error(codes.Unauthenticated, "Your session data is invalid")
	}
	tokenBytes = tokenBytes[:n]
	decrypted, err := p.decryptSession(tokenBytes)
	if err != nil {
		p.logger.Debug("Failed to unmarshal session data", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, "Your session data is invalid")
	}
	session := &providers.SessionInfo{}
	err = msgpack.Unmarshal(decrypted, &session)
	if err != nil {
		p.logger.Debug("Failed to unmarshal session data", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, "Your session data is invalid")
	}

	exp := time.Unix(session.Deadline, 0).UTC()
	if time.Now().UTC().Compare(exp) == -1 {
		return nil, status.Error(codes.Unauthenticated, "Token is invalid or expired")
	}

	env := plugins.Env{
		Method:  method,
		Session: *session,
	}

	check, err := p.plugins.Check(method, env)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You cannot perform this action!")
	}

	return context.WithValue(ctx, "session", session), nil
}

func (p *LocalProvider) InitNewDB() error {
	_, err := p.db.Exec(SetupTables)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}
	// Add root user
	root := memkv.NewMemKV(".", nil)
	// set random root password
	password := p.conf.GetStringOrDefault("root_password", "")
	if password == "" {
		return errors.New("environment variable ROOT_PASSWORD is not set")
	}
	pwKey := kdf.PbKdf2.Key([]byte(password), 4096, 32, sha512.New)
	root.Set("auth.auth_token", pwKey.String())
	root.Set("auth.auth_type", "password")
	root.Set("auth.auth_realm", "local")
	data, err := json.Marshal(root.GetSerializableMap())
	if err != nil {
		return fmt.Errorf("failed to marshal root user info: %w", err)
	}
	_, err = p.db.Exec(AddUserQuery, "root", "root", "local", data)
	if err != nil {
		return fmt.Errorf("failed to add root user: %w", err)
	}
	return nil
}

func (p *LocalProvider) encryptSession(data []byte) ([]byte, error) {
	secureAes, err := encryption.NewSecureAES(p.ephemeralKey, encryption.AES256)
	if err != nil {
		return nil, err
	}

	encrypted, err := secureAes.EncryptToBytes(data)
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

func (p *LocalProvider) decryptSession(encrypted []byte) ([]byte, error) {
	secureAes, err := encryption.NewSecureAES(p.ephemeralKey, encryption.AES256)
	if err != nil {
		return nil, err
	}

	decrypted, err := secureAes.DecryptFromBytes(encrypted)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
