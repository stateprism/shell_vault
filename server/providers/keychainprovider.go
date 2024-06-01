package providers

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"errors"
	"strings"
	"time"

	"github.com/stateprism/prisma_ca/server/lib"
)

type KeyType int

const (
	KEY_TYPE_UNKNOWN   KeyType = iota
	KEY_TYPE_ECDSA_256 KeyType = iota
	KEY_TYPE_ECDSA_384 KeyType = iota
	KEY_TYPE_ECDSA_521 KeyType = iota
	KEY_TYPE_ED25519   KeyType = iota
	KEY_TYPE_RSA_2048  KeyType = iota
	KEY_TYPE_RSA_4096  KeyType = iota
)

var KTStringMap = map[KeyType]string{
	KEY_TYPE_RSA_2048:  "RSA2048",
	KEY_TYPE_RSA_4096:  "RSA4096",
	KEY_TYPE_ECDSA_256: "ECDSA256",
	KEY_TYPE_ECDSA_384: "ECDSA384",
	KEY_TYPE_ECDSA_521: "ECDSA521",
	KEY_TYPE_ED25519:   "ED25519",
	KEY_TYPE_UNKNOWN:   "Unknown",
}
var KTMap = lib.InvertMap(KTStringMap)

func (kt KeyType) String() string {
	return KTStringMap[kt]
}

func KTFromString(s string) KeyType {
	s = strings.ToUpper(s)
	if kt, ok := KTMap[s]; ok {
		return kt
	}
	return KEY_TYPE_UNKNOWN
}

func KTStringArrayToKTArray(kt []string) ([]KeyType, bool, int) {
	ret := make([]KeyType, 0)
	for i, k := range kt {
		k = strings.ToUpper(k)
		// if just 'ECDSA' is specified, add all ECDSA types
		if k == "ECDSA" {
			ret = append(ret, KEY_TYPE_ECDSA_256, KEY_TYPE_ECDSA_384, KEY_TYPE_ECDSA_521)
		} else if k == "RSA" {
			ret = append(ret, KEY_TYPE_RSA_2048, KEY_TYPE_RSA_4096)
		} else {
			kt := KTFromString(k)
			if kt == KEY_TYPE_UNKNOWN {
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
	CRYPTO_ERR_NO__KEY                   CryptoError = iota
	CRYPTO_ERR_NOT_A__KEY                CryptoError = iota
	CRYPTO_ERR_NO_CERTIFICATE            CryptoError = iota
	CRYPTO_ERR_CERTIFICATE_EXPIRED       CryptoError = iota
	CRYPTO_ERR_CERTIFICATE_NOT_YET_VALID CryptoError = iota
	CRYPTO_ERR_CERTIFICATE_CHAIN         CryptoError = iota
	CRYPTO_ERR_CERTIFICATE_INVALID       CryptoError = iota
	CRYPTO_ERR__KEY_TYPE                 CryptoError = iota
)

func (e CryptoError) Error() string {
	switch {
	case errors.Is(e, CRYPTO_ERR_NO__KEY):
		return "No  key available"
	case errors.Is(e, CRYPTO_ERR_NOT_A__KEY):
		return "Data is not a  key"
	case errors.Is(e, CRYPTO_ERR_NO_CERTIFICATE):
		return "No certificate available"
	case errors.Is(e, CRYPTO_ERR_CERTIFICATE_EXPIRED):
		return "Certificate has expired"
	case errors.Is(e, CRYPTO_ERR_CERTIFICATE_NOT_YET_VALID):
		return "Certificate is not yet valid"
	case errors.Is(e, CRYPTO_ERR_CERTIFICATE_CHAIN):
		return "Certificate chain is invalid"
	case errors.Is(e, CRYPTO_ERR_CERTIFICATE_INVALID):
		return "Certificate is invalid"
	case errors.Is(e, CRYPTO_ERR__KEY_TYPE):
		return " key type is not acceptable according to the options provided"
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

type PrivateKey struct {
	id               KeyIdentifier
	ttl              time.Duration
	kType            KeyType
	encrypted        bool
	ServerConstraint string
	key              crypto.PrivateKey
}

// GetKey will return one of the keys corresponding to the key stored in this PrivateKey and nil on others
func (p *PrivateKey) GetKey() (*ecdsa.PrivateKey, *rsa.PrivateKey, *ed25519.PrivateKey) {
	switch p.kType {
	case KEY_TYPE_ED25519:
		k := p.key.(ed25519.PrivateKey)
		return nil, nil, &k
	case KEY_TYPE_ECDSA_256, KEY_TYPE_ECDSA_384, KEY_TYPE_ECDSA_521:
		k := p.key.(ecdsa.PrivateKey)
		return &k, nil, nil
	case KEY_TYPE_RSA_2048, KEY_TYPE_RSA_4096:
		k := p.key.(rsa.PrivateKey)
		return nil, &k, nil
	default:
		return nil, nil, nil
	}
}

func (p *PrivateKey) UnwrapKey() crypto.PrivateKey {
	switch p.kType {
	case KEY_TYPE_ED25519:
		k := p.key.(ed25519.PrivateKey)
		return &k
	case KEY_TYPE_ECDSA_256, KEY_TYPE_ECDSA_384, KEY_TYPE_ECDSA_521:
		k := p.key.(ecdsa.PrivateKey)
		return &k
	case KEY_TYPE_RSA_2048, KEY_TYPE_RSA_4096:
		k := p.key.(rsa.PrivateKey)
		return &k
	default:
		panic("unknown key type on unwrap")
	}
}

func NewPrivateKey(id KeyIdentifier, kt KeyType, key crypto.PrivateKey, ttl time.Duration) *PrivateKey {
	p := &PrivateKey{}
	p.id = id
	p.kType = kt
	p.key = key
	p.ttl = ttl
	return p
}

func (p *PrivateKey) GetTtl() time.Duration {
	return p.ttl
}

func (p *PrivateKey) GetType() KeyType {
	return p.kType
}

type KeyIdentifier interface{}
type KeyLookupCriteria map[string]interface{}
type ExpKeyHook func(p KeychainProvider, identifiers []KeyIdentifier)

type KeychainProvider interface {
	// Unseal is called when ready to use the provider and should decrypt the keys stored
	// can be a no-op by returning true always
	Unseal(key []byte) bool
	// Seal is called before teardown of the server and signals for the provider to do any housekeeping needed
	// can be a no-op
	Seal() bool
	// MakeNewKey create a new  key using KT algorithm
	// a KeyIdentifier which is an opaque type shall be returned
	MakeNewKey(kt KeyType, ttl int64) (KeyIdentifier, error)
	SetActiveKey(kid KeyIdentifier) bool
	GetActiveKey() (KeyIdentifier, bool)
	RetrieveKey(kid KeyIdentifier) (*PrivateKey, bool)
	RetrieveActiveKey() *PrivateKey
	// LookupKey should take a KeyLookupCriteria and return the appropriate
	LookupKey(criteria KeyLookupCriteria) (KeyIdentifier, bool)
	// IsCurrentKey will be checked by the ca server before executing the key rollover by ttl process
	IsCurrentKey(kid KeyIdentifier) bool
	SetExpKeyHook(f ExpKeyHook) ExpKeyHook
	String() string
}
