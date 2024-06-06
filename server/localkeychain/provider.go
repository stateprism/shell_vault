package localkeychain

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/afero"
	"github.com/stateprism/libprisma/cryptoutil"
	"github.com/stateprism/libprisma/cryptoutil/encryption"
	"github.com/stateprism/prisma_ca/server/providers"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

// LocalKeychain must implement providers.KeychainProvider
type LocalKeychain struct {
	fs           afero.Fs
	expHook      providers.ExpKeyHook
	logger       *zap.Logger
	ticker       *time.Ticker
	tickInterval time.Duration
	tickStop     chan struct{}
	db           *sql.DB
	eKey         []byte
}

type LKParams struct {
	fx.In
	Lc     fx.Lifecycle
	Config providers.ConfigurationProvider
	Logger *zap.Logger
	Env    *providers.EnvProvider
}

func NewLocalKeychain(par LKParams) (providers.KeychainProvider, error) {
	// setup key encryption
	if par.Env.GetEnvOrDefault("KEK", "") == "" {
		return nil, errors.New("local keychain provider requires a key encryption key")
	}
	kek := par.Env.GetEnvOrDefault("KEK", "")
	// clear the KEK from the environment
	par.Env.SetEnv("KEK", "")

	// general initialization
	kcPath, err := par.Config.GetString("providers.local_keychain_provider.path")
	if err != nil {
		return nil, err
	}
	kcFsType, err := par.Config.GetString("providers.local_keychain_provider.fs")
	if err != nil {
		kcFsType = "local"
	}

	var fs afero.Fs
	switch kcFsType {
	case "local", "":
		fs = afero.NewOsFs()
		err := os.Chdir(par.Config.GetLocalStore())
		if err != nil {
			return nil, err
		}
		kcPath, err = filepath.Abs(kcPath)
	case "memory":
		fs = afero.NewMemMapFs()
		_ = fs.MkdirAll(kcPath, 0700)
	default:
		return nil, fmt.Errorf("unknown fs type: %s for provider", kcFsType)
	}

	if err != nil {
		return nil, err
	}

	stat, err := fs.Stat(kcPath)
	if os.IsNotExist(err) {
		errDir := fs.MkdirAll(kcPath, 0755)
		if errDir != nil {
			return nil, err
		}
		stat, _ = fs.Stat(kcPath)
	} else if err != nil {
		return nil, err
	}

	if !stat.IsDir() {
		return nil, fmt.Errorf("path %s is a file, this provider requires a directory", kcPath)
	}

	db, err := sql.Open("sqlite3", filepath.Join(kcPath, "keys.db"))
	if err != nil {
		return nil, err
	}
	lk := &LocalKeychain{
		fs:     fs,
		logger: par.Logger,
		db:     db,
		eKey:   cryptoutil.SeededRandomData([]byte(kek), 32),
	}

	if _, err := db.Exec(KeychainExists); err != nil {
		err := lk.initDB()
		if err != nil {
			return nil, err
		}
	}

	// Start the ttl ticker, we check for expired every minute
	if tr, err := par.Config.GetInt64("providers.local_keychain_provider.ttl_tick"); err != nil {
		lk.logger.Warn("TTL tick is not configured setting to 60s default")
		lk.tickInterval = 60 * time.Second
	} else {
		lk.logger.Info("TTL tick will be configured to", zap.Duration("tickRate", time.Duration(tr)*time.Second))
		lk.tickInterval = time.Duration(tr) * time.Second
	}

	par.Lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lk.ticker = time.NewTicker(lk.tickInterval)
			lk.tickStop = make(chan struct{})
			go ttlTick(lk)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			close(lk.tickStop)
			err := lk.db.Close()
			if err != nil {
				return err
			}
			return nil
		},
	})

	return lk, nil
}

func ttlTick(l *LocalKeychain) {
	for {
		select {
		case <-l.ticker.C:
			now := time.Now()
			toDrop := make([]providers.KeyIdentifier, 0)
			l.logger.Debug(
				"Checking ttl of stored keys:",
				zap.Time("event_at", now),
				zap.Time("next_check", now.Add(60*time.Second)),
			)
			rows, err := l.db.Query(SelectExpiredKeys, now.UTC().Unix())
			if err != nil {
				l.logger.Error("failed to query for expired keys", zap.Error(err))
				continue
			}
			for rows.Next() {
				var keyName string
				err := rows.Scan(&keyName)
				if err != nil {
					l.logger.Error("failed to scan key name", zap.Error(err))
					continue
				}
				l.logger.Info("Key expired", zap.String("key_name", keyName))
				toDrop = append(toDrop, keyName)
			}
			err = rows.Close()
			if err != nil {
				panic(err)
			}
			for _, k := range toDrop {
				l.DropKey(k)
			}
			if l.expHook == nil {
				l.logger.Warn("No hook set on ttl provider to notify ca")
				continue
			}
			l.expHook(l, toDrop)
		case <-l.tickStop:
			l.logger.Info("stopping ttl ticks")
			l.ticker.Stop()
			return
		}
	}
}

func (l *LocalKeychain) Unseal(key []byte) bool {
	return true
}

func (l *LocalKeychain) Seal() bool {
	err := l.db.Close()
	if err != nil {
		return false
	}
	return true
}

func makeNewEcdsaKey(kt providers.KeyType) (ecdsa.PrivateKey, error) {
	var c elliptic.Curve
	switch kt {
	case providers.KEY_TYPE_ECDSA_256:
		c = elliptic.P256()
	case providers.KEY_TYPE_ECDSA_384:
		c = elliptic.P384()
	case providers.KEY_TYPE_ECDSA_521:
		c = elliptic.P521()
	default:
		panic("unhandled default case")
	}
	t, err := ecdsa.GenerateKey(c, cryptorand.Reader)
	if err != nil {
		return ecdsa.PrivateKey{}, err
	}
	return *t, nil
}

func makeRsaKey(kt providers.KeyType) (rsa.PrivateKey, error) {
	var t *rsa.PrivateKey
	var bits int
	switch kt {
	case providers.KEY_TYPE_RSA_2048:
		bits = 2048
	case providers.KEY_TYPE_RSA_4096:
		bits = 4096
	default:
		panic("unhandled default case")
	}
	t, err := rsa.GenerateKey(cryptorand.Reader, bits)
	if err != nil {
		return rsa.PrivateKey{}, err
	}
	return *t, nil
}

func (l *LocalKeychain) MakeNewKey(keyName providers.KeyIdentifier, kt providers.KeyType, ttl int64) (providers.KeyIdentifier, error) {
	if _, ok := keyName.(string); !ok || keyName == nil {
		return nil, errors.New("this provider only takes string key names")
	}

	var key crypto.PrivateKey
	var encrypted []byte
	var err error
	switch kt {
	case providers.KEY_TYPE_ED25519:
		var t ed25519.PrivateKey
		_, t, err = ed25519.GenerateKey(cryptorand.Reader)
		key = t
	case providers.KEY_TYPE_ECDSA_256, providers.KEY_TYPE_ECDSA_384, providers.KEY_TYPE_ECDSA_521:
		key, err = makeNewEcdsaKey(kt)
	case providers.KEY_TYPE_RSA_2048, providers.KEY_TYPE_RSA_4096:
		key, err = makeRsaKey(kt)
	default:
		return nil, fmt.Errorf("invalid key format: %s", kt)
	}
	if err != nil {
		return nil, err
	}

	providers.NewPrivateKey(keyName, kt, key, time.Duration(ttl))
	enc, err := encryption.NewSecureAES(l.eKey)
	if err != nil {
		return nil, err
	}
	encrypted, err = enc.EncryptToBytes(key.(ed25519.PrivateKey).Seed())
	if err != nil {
		return nil, err
	}

	validTo := time.Now().UTC().Add(time.Duration(ttl) * time.Second)
	_, err = l.db.Exec(InsertKey, keyName, kt.String(), encrypted, validTo, ttl)
	if err != nil {
		return nil, err
	}

	return keyName, nil
}

func (l *LocalKeychain) SetExpKeyHook(f providers.ExpKeyHook) providers.ExpKeyHook {
	old := l.expHook
	l.expHook = f
	return old
}

func (l *LocalKeychain) LookupKey(criteria providers.KeyLookupCriteria) (providers.KeyIdentifier, bool) {
	//TODO implement me
	panic("implement me")
}

func (l *LocalKeychain) RetrieveKey(kid providers.KeyIdentifier) (*providers.PrivateKey, error) {
	_, ok := kid.(string)
	if !ok {
		return nil, errors.New("this provider only takes string key names")
	}
	encryptedKey := make([]byte, 64)
	row := l.db.QueryRow(SelectKey, kid)
	var keyType string
	var ttl int64
	err := row.Scan(&encryptedKey, &keyType, &ttl)
	if err != nil {
		return nil, err
	}
	enc, err := encryption.NewSecureAES(l.eKey)
	if err != nil {
		return nil, err
	}
	key, err := enc.DecryptFromBytes(encryptedKey)
	if err != nil {
		return nil, err
	}
	switch providers.KTFromString(keyType) {
	case providers.KEY_TYPE_ED25519:
		k := ed25519.NewKeyFromSeed(key)
		return providers.NewPrivateKey(kid, providers.KEY_TYPE_ED25519, k, time.Duration(ttl)), nil
	default:
		panic("unhandled default case")
	}
}

func (l *LocalKeychain) DropKey(keyName providers.KeyIdentifier) bool {
	if _, ok := keyName.(string); !ok {
		return false
	}
	_, err := l.db.Exec(DropKey, keyName)
	if err != nil {
		l.logger.Error("failed to delete key", zap.Error(err))
		return false
	}
	return true
}

func (l *LocalKeychain) MakeAndReplaceKey(keyName providers.KeyIdentifier, kt providers.KeyType, ttl int64) (providers.KeyIdentifier, error) {
	if _, ok := keyName.(string); !ok || keyName == nil {
		return nil, errors.New("this provider only takes string key names")
	}
	l.DropKey(keyName.(string))

	return l.MakeNewKey(keyName, kt, ttl)
}

func (l *LocalKeychain) MakeNewKeyIfNotExists(keyName providers.KeyIdentifier, kt providers.KeyType, ttl int64) (providers.KeyIdentifier, error) {
	if _, ok := keyName.(string); !ok || keyName == nil {
		return nil, errors.New("this provider only takes string key names")
	}
	if _, err := l.RetrieveKey(keyName); err != nil {
		return keyName, nil
	}
	return l.MakeNewKey(keyName, kt, ttl)
}

func (l *LocalKeychain) String() string {
	return "LocalKeychain"
}

func (l *LocalKeychain) saveKey(keyName string) {

}

func (l *LocalKeychain) initDB() error {
	_, err := l.db.Exec(CreateTables)
	if err != nil {
		l.logger.Fatal("failed to create key table", zap.Error(err))
		return err
	}
	return nil
}
