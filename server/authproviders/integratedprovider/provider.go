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
	"github.com/stateprism/prisma_ca/server/middleware"
	"github.com/stateprism/prisma_ca/server/providers"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"path"
	"time"
)

type LocalProvider struct {
	conf         providers.ConfigurationProvider
	kv           *memkv.MemKV
	ephemeralKey []byte
	logger       *zap.Logger
	db           *sql.DB
	env          *providers.EnvProvider
}

func New(config providers.ConfigurationProvider, kv *memkv.MemKV, log *zap.Logger) (providers.AuthProvider, error) {
	key := kdf.PbKdf2.Key(cryptoutil.NewRandom(128), 4096, 32, sha512.New)
	localStore := config.GetLocalStore()
	if localStore == "" {
		return nil, fmt.Errorf("local store path not found")
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
		kv:           kv,
		ephemeralKey: key.GetKey(),
		logger:       log,
		env:          providers.NewEnvProvider("SHELL_VAULT_"),
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
	return "PAM"
}

func (p *LocalProvider) Authenticate(ctx context.Context) (string, error) {
	token, err := middleware.AuthFromMetadata(ctx, "local", "authorization")
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
	if pwKey.Equals(string(password)) {
		time.Sleep(1 * time.Second)
		return "", status.Error(codes.Unauthenticated, "Invalid username or password")
	}

	sesId := uuid.New()
	sesInfo := memkv.NewMemKV(".", nil)
	sesInfo.Set("session.user.username", string(username))
	sesInfo.Set("session.user.realm", "local")
	sesInfo.Set("session.id", sesId)
	sesInfo.Set("session.expires", time.Now().Add(20*time.Hour))
	p.kv.Set("sessions."+sesId.String(), entInfo)
	data, err = json.Marshal(sesInfo.GetSerializableMap())
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

func (p *LocalProvider) GetSession(ctx context.Context) (context.Context, error) {
	token, err := middleware.AuthFromMetadata(ctx, "encrypted", "authorization")
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
	tokenData := make(map[string]any)
	err = json.Unmarshal(decrypted, &tokenData)
	if err != nil {
		p.logger.Debug("Failed to unmarshal session data", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, "Your session data is invalid")
	}

	entInfo := memkv.NewMemKV(".", nil)
	err = entInfo.LoadFromSerializableMap(tokenData)
	if err != nil {
		p.logger.Debug("Failed to unmarshal session data", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, "Your session data is invalid")
	}

	return context.WithValue(ctx, "session", entInfo), nil
}

func (p *LocalProvider) InitNewDB() error {
	_, err := p.db.Exec(SetupTables)
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}
	// Add root user
	root := memkv.NewMemKV(".", nil)
	// set random root password
	password := p.env.GetEnvOrDefault("ROOT_PASSWORD", "")
	if password == "" {
		return errors.New("environment variable ROOT_PASSWORD is not set")
	}
	// clear the password from the environment
	p.env.SetEnv("ROOT_PASSWORD", "")
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
	secureAes, err := encryption.NewSecureAESWithSafeKey(p.ephemeralKey)
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
	secureAes, err := encryption.NewSecureAESWithSafeKey(p.ephemeralKey)
	if err != nil {
		return nil, err
	}

	decrypted, err := secureAes.DecryptFromBytes(encrypted)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
