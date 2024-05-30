package providers

import (
	"crypto"
	"slices"
	"strings"

	"github.com/stateprism/prisma_ca/server/lib"
)

type KeyType int

const (
	PRIVATEKEY_TYPE_UNKNOWN   KeyType = iota
	PRIVATEKEY_TYPE_RSA       KeyType = iota
	PRIVATEKEY_TYPE_ECDSA_256 KeyType = iota
	PRIVATEKEY_TYPE_ECDSA_384 KeyType = iota
	PRIVATEKEY_TYPE_ECDSA_521 KeyType = iota
	PRIVATEKEY_TYPE_ED25519   KeyType = iota
)

var KTStringMap = map[KeyType]string{
	PRIVATEKEY_TYPE_RSA:       "RSA",
	PRIVATEKEY_TYPE_ECDSA_256: "ECDSA256",
	PRIVATEKEY_TYPE_ECDSA_384: "ECDSA384",
	PRIVATEKEY_TYPE_ECDSA_521: "ECDSA521",
	PRIVATEKEY_TYPE_ED25519:   "ED25519",
	PRIVATEKEY_TYPE_UNKNOWN:   "Unknown",
}
var KTMap = lib.InvertMap(KTStringMap)

func (kt KeyType) String() string {
	return KTStringMap[kt]
}

func KTFromString(s string) KeyType {
	if kt, ok := KTMap[s]; ok {
		return kt
	}
	return PRIVATEKEY_TYPE_UNKNOWN
}

func KTStringArrayToKTArray(kt []string) ([]KeyType, bool, int) {
	ret := make([]KeyType, 0)
	for i, k := range kt {
		k = strings.ToUpper(k)
		// if just 'ECDSA' is specified, add all ECDSA types
		if k == "ECDSA" {
			ret = append(ret, PRIVATEKEY_TYPE_ECDSA_256, PRIVATEKEY_TYPE_ECDSA_384, PRIVATEKEY_TYPE_ECDSA_521)
		} else {
			kt := KTFromString(k)
			if kt == PRIVATEKEY_TYPE_UNKNOWN {
				return nil, false, i
			}
			ret = append(ret, kt)
		}
	}
	return ret, true, -1
}

func KTArrayToKTStringArray(kt []KeyType) []string {
	ret := make([]string, 0)
	for _, k := range kt {
		ret = append(ret, k.String())
	}
	return ret
}

type CryptoError int

const (
	CRYPTO_ERR_UNKNOWN                   CryptoError = iota
	CRYPTO_ERR_NO_PRIVATE_KEY            CryptoError = iota
	CRYPTO_ERR_NOT_A_PRIVATE_KEY         CryptoError = iota
	CRYPTO_ERR_NO_CERTIFICATE            CryptoError = iota
	CRYPTO_ERR_CERTIFICATE_EXPIRED       CryptoError = iota
	CRYPTO_ERR_CERTIFICATE_NOT_YET_VALID CryptoError = iota
	CRYPTO_ERR_CERTIFICATE_CHAIN         CryptoError = iota
	CRYPTO_ERR_CERTIFICATE_INVALID       CryptoError = iota
	CRYPTO_ERR_PRIVATE_KEY_TYPE          CryptoError = iota
)

func (e CryptoError) Error() string {
	switch e {
	case CRYPTO_ERR_NO_PRIVATE_KEY:
		return "No private key available"
	case CRYPTO_ERR_NOT_A_PRIVATE_KEY:
		return "Data is not a private key"
	case CRYPTO_ERR_NO_CERTIFICATE:
		return "No certificate available"
	case CRYPTO_ERR_CERTIFICATE_EXPIRED:
		return "Certificate has expired"
	case CRYPTO_ERR_CERTIFICATE_NOT_YET_VALID:
		return "Certificate is not yet valid"
	case CRYPTO_ERR_CERTIFICATE_CHAIN:
		return "Certificate chain is invalid"
	case CRYPTO_ERR_CERTIFICATE_INVALID:
		return "Certificate is invalid"
	case CRYPTO_ERR_PRIVATE_KEY_TYPE:
		return "Private key type is not acceptable according to the options provided"
	default:
		return "Unknown error"
	}
}

type CryptoLoaderOptions struct {
	// Certificate loader options
	AllowExpiredCertificates bool
	VerifyCertChain          bool
	// A list of acceptable key types, if not set all key types are rejected
	AcceptableKeyTypes []KeyType
}

type CryptoProvider interface {
	// ProviderIsBlackBox returns true if the provider does not return a private key, but instead does black box signing
	ProviderIsBlackBox() bool
	// LoadCertificate loads a certificate from the provider and returns a SignatureProvider
	LoadCertificate(certName string) (*SignatureProvider, error)
	// LoadPrivateKey loads a private key from the provider and returns a SignatureProvider
	LoadPrivateKey(keyName, password string) (*SignatureProvider, error)
	// NewAuthority create a new certificate authority with the given options
	NewAuthority(name, password, comment string, kt KeyType) error
	// RemoteSignPayloadWithKey is used to sign a payload with a key that is not stored locally, when the provider won't expose a private key
	RemoteSignPayloadWithKey(keyIdentifier interface{}, payload []byte) ([]byte, error)
	// RemoteVerifySignatureWithKey is used to verify a signature with a key that is not stored locally
	RemoteVerifySignatureWithKey(keyIdentifier interface{}, payload []byte, signature []byte) (bool, error)
}

type SignaturePrivateKey struct {
	Material          *crypto.PrivateKey
	EncryptedMaterial []byte
	IsEncrypted       bool
	Type              KeyType
}

type SignaturePublicKey struct {
	Material *crypto.PublicKey
	Type     KeyType
}

type SignatureCryptoMaterial struct {
	Private *SignaturePrivateKey
	Pub     *crypto.PublicKey
}

type SignatureProvider struct {
	Opts     *CryptoLoaderOptions
	Material *SignatureCryptoMaterial
}

func (p *SignatureProvider) SignPayload(payload []byte) ([]byte, error) {
	if p.Material.Private == nil {
		return nil, CRYPTO_ERR_NO_PRIVATE_KEY
	}

	if len(p.Opts.AcceptableKeyTypes) == 0 || slices.Contains(p.Opts.AcceptableKeyTypes, p.Material.Private.Type) {
		return nil, CRYPTO_ERR_PRIVATE_KEY_TYPE
	}

	return nil, nil
}

func (p *SignatureProvider) VerifySignature(payload []byte, signature []byte) (bool, error) {
	return false, nil
}
