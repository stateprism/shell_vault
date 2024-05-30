package plainfileprovider

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/afero"
	pb "github.com/stateprism/prisma_ca/rpc/caproto"
	"github.com/stateprism/prisma_ca/server/authproviders"
)

type SyncMap struct {
	l sync.RWMutex
	m map[string]string
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		l: sync.RWMutex{},
		m: make(map[string]string),
	}
}

func (sm *SyncMap) Get(key string) (string, bool) {
	sm.l.RLock()
	defer sm.l.RUnlock()
	v, ok := sm.m[key]
	return v, ok
}

func (sm *SyncMap) Set(key, value string) {
	sm.l.Lock()
	defer sm.l.Unlock()
	sm.m[key] = value
}

type PlainFileProvider struct {
	Users *SyncMap
}

type pfp struct {
	Entries map[string]string `json:"entries"`
}

func New(fs afero.Fs, filename string) (*PlainFileProvider, error) {
	if fs == nil {
		return nil, fmt.Errorf("fs is nil")
	}
	if filename == "" {
		return nil, fmt.Errorf("filename is empty")
	}

	if _, err := os.Stat(filename); err != nil {
		return nil, fmt.Errorf("file does not exist")
	}

	fd, err := fs.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	stat, _ := fd.Stat()
	buf := make([]byte, stat.Size())

	if _, err := fd.Read(buf); err != nil {
		return nil, err
	}

	p := &pfp{}
	if err := json.Unmarshal(buf, p); err != nil {
		return nil, err
	}

	sm := NewSyncMap()
	for k, v := range p.Entries {
		sm.Set(k, v)
	}

	return &PlainFileProvider{
		Users: sm,
	}, nil
}

func (p *PlainFileProvider) String() string {
	return "plainfile"
}

func (p *PlainFileProvider) Authenticate(ctx context.Context, msg *pb.AuthRequest) (bool, error) {
	req, err := authproviders.RequestFromBytes(msg.GetAuthRequest())
	if err != nil {
		return false, fmt.Errorf("invalid request")
	}

	pw, ok := p.Users.Get(req.Username)
	if !ok {
		return false, nil
	}

	return pw == req.Password, nil
}

func (p *PlainFileProvider) GetUserIdentifier(ctx context.Context, msg *pb.AuthRequest) string {
	req, _ := authproviders.RequestFromBytes(msg.GetAuthRequest())
	return req.Username
}
