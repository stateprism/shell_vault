package certificate

import (
	"crypto/sha256"
	"fmt"
	"github.com/spf13/afero"
	"google.golang.org/protobuf/proto"
	"slices"
)

func (s *AuthorizedUsers) CalculateIntegrityHash() [32]byte {
	sCopy := proto.Clone(s).(*AuthorizedUsers)
	sCopy.IntegrityHash = nil
	cBytes, _ := proto.Marshal(sCopy)
	hash := sha256.Sum256(cBytes)
	return hash
}

func (s *AuthorizedUsers) VerifyIntegrityHash() bool {
	hash := s.CalculateIntegrityHash()
	return slices.Compare(hash[:], s.GetIntegrityHash()) == 0
}

func (s *AuthorizedUsers) UpdateIntegrityHash() {
	hash := s.CalculateIntegrityHash()
	s.IntegrityHash = hash[:]
}

func (s *AuthorizedUsers) SaveUserStore(path string) error {
	s.UpdateIntegrityHash()
	fs := afero.NewOsFs()
	contents, err := proto.Marshal(s)
	if err != nil {
		return fmt.Errorf("failed to marshal authorized users: %v", err)
	}

	f, err := fs.Create(path)
	if err != nil {
		return fmt.Errorf("failed to open authorized users file: %v", err)
	}

	defer f.Close()

	f.Write(contents)
	return nil
}
