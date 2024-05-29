package certificate

import (
	"crypto/sha256"
	"fmt"
	"slices"
	"time"

	"github.com/spf13/afero"
	"google.golang.org/protobuf/proto"
)

func LoadUserStore(path string, createIfNotExists bool) (*AuthorizedUsers, error) {
	fs := afero.NewOsFs()
	stat, err := fs.Stat(path)
	if err != nil && !createIfNotExists {
		return nil, fmt.Errorf("failed to read authorized users file: %v", err)
	} else if err != nil && createIfNotExists {
		store := &AuthorizedUsers{
			Users: make(map[string]*AuthorizedUser),
		}
		return store, nil
	}

	contents := make([]byte, stat.Size())
	userStore := &AuthorizedUsers{}

	f, err := fs.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open authorized users file: %v", err)
	}

	defer f.Close()

	f.Read(contents)
	if err := proto.Unmarshal(contents, userStore); err != nil {
		return nil, fmt.Errorf("failed to unmarshal authorized users: %v", err)
	}

	return userStore, nil
}

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

func (s *AuthorizedUsers) AddUser(username string, key string) {
	s.Users[username] = &AuthorizedUser{
		UserName:      username,
		UserPublicKey: key,
		CreatedAt:     time.Now().Unix(),
	}
}

func (s *AuthorizedUsers) RemoveUser(username string) {
	s.Users[username] = nil
}

func (s *AuthorizedUsers) GetUser(username string) (*AuthorizedUser, bool) {
	user, ok := s.Users[username]
	return user, ok
}
