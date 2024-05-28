package certhandler

import (
	"crypto"
	"crypto/x509"
	"time"

	"github.com/google/uuid"
)

type CertUsage int

const (
	CERT_USAGE_HOST_VERIFY CertUsage = 0
	CERT_USAGE_USER_AUTH   CertUsage = 1
)

type Cert struct {
	verified   bool
	cert       x509.Certificate
	publicKey  crypto.PublicKey
	privateKey crypto.PrivateKey
}

type CertProvider interface {
	FindCertByID(id uuid.UUID) Cert
	FindCertByAttribute(attr string, val interface{}) Cert
	GetIntermediateChain() *x509.CertPool
	GetRoots() *x509.CertPool
	FindPrivateKeyForCert(cert Cert)
}

func NewFromX509(cert x509.Certificate, verify bool, p CertProvider) (*Cert, error) {
	new := &Cert{
		verified:   false,
		cert:       cert,
		publicKey:  cert.PublicKey,
		privateKey: nil,
	}

	if verify {
		opts := x509.VerifyOptions{
			Intermediates: p.GetIntermediateChain(),
			Roots:         p.GetRoots(),
			CurrentTime:   time.Now(),
		}

		_, err := cert.Verify(opts)
		if err != nil {
			return nil, err
		}

		new.verified = true
	}

	return new, nil
}
