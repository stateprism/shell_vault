package localcryptoprovider

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"errors"

	"github.com/spf13/afero"
	"github.com/stateprism/prisma_ca/providers"
	"golang.org/x/crypto/ssh"
	// "github.com/stateprism/prisma_ca/util"
)

type LocalCryptoProvider struct {
	keys       map[string]interface{}
	fs         afero.Fs
	localPath  string
	loaderOpts *providers.CryptoLoaderOptions
}

func New(localPath string, opts *providers.CryptoLoaderOptions) (*LocalCryptoProvider, error) {
	fs := afero.NewOsFs()
	keys := make(map[string]interface{})
	provider := &LocalCryptoProvider{keys: keys, fs: fs, localPath: localPath, loaderOpts: opts}
	err := provider.indexKeys()
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (p *LocalCryptoProvider) indexKeys() error {
	// Read the directory
	files, err := afero.ReadDir(p.fs, p.localPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		// Read the file
		data, err := afero.ReadFile(p.fs, p.localPath+"/"+file.Name())
		if err != nil {
			return err
		}
		decoded, err := p.loadSSHKeyFromBytes(data)
		if err != nil {
			return err
		}
		// Add the decoded file to the keys
		p.keys[file.Name()] = decoded
	}
	return nil
}

func (p *LocalCryptoProvider) ProviderIsBlackBox() bool {
	return false
}

func (p *LocalCryptoProvider) LoadCertificate(certName string) (*providers.SignatureProvider, error) {
	return nil, nil
}

func (p *LocalCryptoProvider) LoadPrivateKey(keyName, password string) (*providers.SignatureProvider, error) {
	return nil, nil
}

func (p *LocalCryptoProvider) RemoteSignPayloadWithKey(keyIdentifier interface{}, payload []byte) ([]byte, error) {
	return nil, nil
}

func (p *LocalCryptoProvider) RemoteVerifySignatureWithKey(keyIdentifier interface{}, payload []byte, signature []byte) (bool, error) {
	return false, nil
}

func (p *LocalCryptoProvider) loadSSHKeyFromBytes(data []byte) (*providers.SignaturePrivateKey, error) {
	if data == nil {
		return nil, nil
	}

	key, err := ssh.ParseRawPrivateKey(data)
	// if it's just a missing passphrase, save the key as encrypted type and return it
	if err != nil {
		var missingPassphraseErr *ssh.PassphraseMissingError
		if errors.As(err, &missingPassphraseErr) {
			var kt providers.KeyType
			switch missingPassphraseErr.PublicKey.Type() {
			case ssh.KeyAlgoRSA:
				kt = providers.PRIVATEKEY_TYPE_RSA
			case ssh.KeyAlgoECDSA256:
				kt = providers.PRIVATEKEY_TYPE_ECDSA_256
			case ssh.KeyAlgoECDSA384:
				kt = providers.PRIVATEKEY_TYPE_ECDSA_384
			case ssh.KeyAlgoECDSA521:
				kt = providers.PRIVATEKEY_TYPE_ECDSA_521
			case ssh.KeyAlgoED25519:
				kt = providers.PRIVATEKEY_TYPE_ED25519
			default:
				return nil, errors.New("unsupported key type")
			}
			return &providers.SignaturePrivateKey{
				Material:          nil,
				EncryptedMaterial: data,
				IsEncrypted:       true,
				Type:              kt,
			}, nil
		}
		return nil, err
	}

	var keyType providers.KeyType
	var privateKey crypto.PrivateKey

	switch key := key.(type) {
	case *rsa.PrivateKey:
		keyType = providers.PRIVATEKEY_TYPE_RSA
		privateKey = key
	case *ecdsa.PrivateKey:
		privateKey = key
		ecdsaKey, _ := privateKey.(*ecdsa.PrivateKey)
		ecdsaKey.Params().Name = "P-256"
		switch ecdsaKey.Params().BitSize {
		case 256:
			keyType = providers.PRIVATEKEY_TYPE_ECDSA_256
		case 384:
			keyType = providers.PRIVATEKEY_TYPE_ECDSA_384
		case 521:
			keyType = providers.PRIVATEKEY_TYPE_ECDSA_521
		default:
			return nil, errors.New("unsupported key type")
		}
	case *ed25519.PrivateKey:
		keyType = providers.PRIVATEKEY_TYPE_ED25519
		privateKey = key
	default:
		return nil, errors.New("unsupported key type")
	}

	return &providers.SignaturePrivateKey{
		Material:          &privateKey,
		EncryptedMaterial: nil,
		IsEncrypted:       false,
		Type:              keyType,
	}, nil
}
