package servers

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"

	"github.com/google/uuid"
	pb "github.com/stateprism/prisma_ca/rpc/adminproto"
	"github.com/stateprism/prisma_ca/server/providers"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AdminServerParams struct {
	fx.In
	Config providers.ConfigurationProvider
	Env    *providers.EnvProvider
	Logger *zap.Logger
}

type AuthChallenge struct {
	Id        uuid.UUID
	Challenge []byte
	Expires   int64
	Nonce     uint64
}

func NewAuthChallenge() AuthChallenge {
	bytes := make([]byte, 256)
	rand.Read(bytes)
	return AuthChallenge{
		Id:        uuid.New(),
		Challenge: bytes,
		Expires:   time.Now().Add(time.Minute * 5).Unix(),
		Nonce:     uint64(time.Now().Unix()),
	}
}

type ChallengeStore struct {
	Lock       sync.RWMutex
	Challenges map[uuid.UUID]AuthChallenge
}

type AdminServer struct {
	pb.UnsafeAdminServiceServer
	Config     providers.ConfigurationProvider
	Logger     *zap.Logger
	Env        *providers.EnvProvider
	Challenges *ChallengeStore
	Keys       map[string]*ed25519.PublicKey
	Listen     string
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

func NewAdminServer(p AdminServerParams) (*AdminServer, error) {
	s := new(AdminServer)
	s.Config = p.Config
	s.Logger = p.Logger
	s.Env = p.Env

	keys := make(map[string]*ed25519.PublicKey)

	s.Challenges = NewChallengeStore()
	listenAddr, err := s.Config.GetString("admin_server.listen")
	if err != nil {
		return nil, err
	}
	s.Listen = listenAddr
	rootKey, err := s.Config.Get("admin_server.root_key")
	if err != nil {
		return nil, err
	}
	rootKeyBytes, err := base64.RawStdEncoding.DecodeString(rootKey.(string))
	if err != nil {
		return nil, err
	}
	rootKeyPub := ed25519.PublicKey(rootKeyBytes)
	keys["root"] = &rootKeyPub
	s.Keys = keys

	return s, nil
}

func (a *AdminServer) AdminAuthInit(context.Context, *pb.Empty) (*pb.AdminAuthInitResponse, error) {
	challenge := NewAuthChallenge()
	a.Challenges.AddChallenge(challenge)

	return &pb.AdminAuthInitResponse{
		Challenge:      challenge.Challenge,
		ChallengeId:    []byte(challenge.Id.String()),
		ChallengeNonce: challenge.Nonce,
	}, nil
}

func (a *AdminServer) AdminAuthRespond(c context.Context, r *pb.AdminAuthRequest) (*pb.AdminAuthResponse, error) {
	if r.GetChallengeId() == nil {
		return &pb.AdminAuthResponse{
			Success: false,
		}, nil
	}

	return &pb.AdminAuthResponse{
		Success: true,
	}, nil
}

func (a *AdminServer) ChangeRootCert(context.Context, *pb.ChangeRootCertRequest) (*pb.ChangeRootCertResponse, error) {
	panic("unimplemented")
}

func (a *AdminServer) AddUser(ctx context.Context, r *pb.AddUserRequest) (*pb.UserActionResponse, error) {
	return nil, nil
}

func (a *AdminServer) StopServer(context.Context, *pb.StopServerRequest) (*pb.Empty, error) {
	panic("unimplemented")
}

func (a *AdminServer) RegisterServer(s *grpc.Server) {
	pb.RegisterAdminServiceServer(s, a)
}
