package server

import (
	"context"
	"crypto/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	pb "github.com/stateprism/prisma_ca/adminprotocol"
	"github.com/stateprism/prisma_ca/providers"
)

type AuthChallenge struct {
	Id        uuid.UUID
	Challenge []byte
	Expires   int64
}

func NewAuthChallenge() AuthChallenge {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return AuthChallenge{
		Id:        uuid.New(),
		Challenge: bytes,
		Expires:   time.Now().Add(time.Minute * 5).Unix(),
	}
}

type ChallengeStore struct {
	Lock       sync.RWMutex
	Challenges map[uuid.UUID]AuthChallenge
}

type AdminServer struct {
	pb.UnsafeAdminServiceServer
	Config     providers.ConfigurationProvider
	Logger     providers.LogProvider
	Env        *providers.EnvProvider
	Challenges *ChallengeStore
}

func NewChallengeStore() *ChallengeStore {
	return &ChallengeStore{
		Lock:       sync.RWMutex{},
		Challenges: make(map[uuid.UUID]AuthChallenge),
	}
}

func (c *ChallengeStore) AddChallenge(challenge AuthChallenge) {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.Challenges[challenge.Id] = challenge
}

func (c *ChallengeStore) GetChallenge(id uuid.UUID) (AuthChallenge, bool) {
	c.Lock.RLock()
	defer c.Lock.RUnlock()

	challenge, ok := c.Challenges[id]
	return challenge, ok
}

func NewAdminServer(configProvider providers.ConfigurationProvider, envProvider *providers.EnvProvider, log providers.LogProvider) pb.AdminServiceServer {
	server := new(AdminServer)
	server.Config = configProvider
	server.Logger = log
	server.Env = envProvider

	return server
}

// AdminAuthInit implements adminprotocol.AdminServiceServer.
func (a *AdminServer) AdminAuthInit(context.Context, *pb.Empty) (*pb.AdminAuthInitResponse, error) {
	panic("unimplemented")
}

// AdminAuthRespond implements adminprotocol.AdminServiceServer.
func (a *AdminServer) AdminAuthRespond(context.Context, *pb.AdminAuthRequest) (*pb.AdminAuthResponse, error) {
	panic("unimplemented")
}

// ChangeRootCert implements adminprotocol.AdminServiceServer.
func (a *AdminServer) ChangeRootCert(context.Context, *pb.ChangeRootCertRequest) (*pb.ChangeRootCertResponse, error) {
	panic("unimplemented")
}

// StopServer implements adminprotocol.AdminServiceServer.
func (a *AdminServer) StopServer(context.Context, *pb.StopServerRequest) (*pb.Empty, error) {
	panic("unimplemented")
}
