package servers

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"

	pb "github.com/stateprism/prisma_ca/rpc/caproto"
	"github.com/stateprism/prisma_ca/server/lib"
	"github.com/stateprism/prisma_ca/server/providers"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type CaServer struct {
	pb.UnsafePrismaCaServer
	Config    providers.ConfigurationProvider
	Auth      providers.AuthProvider
	DB        providers.DatabaseProvider
	Crypto    providers.CryptoProvider
	AllowedKT []providers.KeyType
	Env       *providers.EnvProvider
	Log       *zap.Logger
	Listen    string
}

type CaServerParams struct {
	fx.In
	Logger *zap.Logger
	Config providers.ConfigurationProvider
	Auth   providers.AuthProvider
	Env    *providers.EnvProvider
}

func NewCAServer(p CaServerParams) (*CaServer, error) {
	s := &CaServer{
		Config: p.Config,
		Auth:   p.Auth,
		Log:    p.Logger,
	}
	if s.Config == nil {
		return nil, fmt.Errorf("config is nil")
	}
	akt, err := s.Config.Get("ca_server.allowed_key_types")
	if err != nil {
		return nil, err
	}
	kt, _ := akt.([]interface{})
	s.AllowedKT, _, _ = providers.KTStringArrayToKTArray(lib.InterfaceArrayToArray[string](kt))
	s.Listen, err = s.Config.GetString("ca_server.listen")
	if err != nil {
		return nil, err
	}

	if s.Auth == nil {
		return nil, fmt.Errorf("auth is nil")
	}

	return s, nil
}

func (s *CaServer) Authenticate(ctx context.Context, msg *pb.AuthRequest) (*pb.AuthReply, error) {
	authTime := uint64(time.Now().Unix())
	authUntil := uint64(time.Now().Add(time.Hour).Unix())
	authSuccess, err := s.Auth.Authenticate(ctx, msg)
	if err != nil {
		s.Log.Error("Error authenticating", zap.Error(err))
		return &pb.AuthReply{
			Success:  false,
			AuthTime: authTime,
			Errors: &pb.Errors{
				Errors: map[string]string{"AuthError": "Internal error"},
			},
		}, nil
	}
	if authSuccess {
		secret, err := s.Config.GetBytes("ca_server.secret")
		if err != nil {
			s.Log.Error("Error getting secret", zap.Error(err))
			return nil, fmt.Errorf("internal error")
		}
		token := generateToken(s.Auth.GetUserIdentifier(ctx, msg), secret)
		return &pb.AuthReply{
			AuthTime:  authTime,
			AuthToken: token,
			AuthUntil: authUntil,
			Success:   true,
		}, nil
	} else {
		return &pb.AuthReply{
			Success:  false,
			AuthTime: authTime,
			Errors: &pb.Errors{
				Errors: map[string]string{"AuthError": "Authentication failed"},
			},
		}, nil
	}
}

func generateToken(username string, b []byte) []byte {
	var token []byte

	time := time.Now().Unix()
	timeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timeBytes, uint64(time))
	token = append(token, timeBytes...)
	token = append(token, []byte(username)...)
	hmac := hmac.New(sha256.New, b)
	hmac.Write(token)
	bytes := hmac.Sum(nil)
	return bytes
}

func (s *CaServer) RequestCert(ctx context.Context, msg *pb.CertRequest) (*pb.CertReply, error) {
	return &pb.CertReply{
		Cert:       []byte("cert"),
		ValidUntil: uint64(time.Now().Add(time.Hour).Unix()),
	}, nil
}

func (s *CaServer) GetConfig(context.Context, *pb.ConfigRequest) (*pb.ConfigReply, error) {
	serverId, _ := s.Config.GetString("server_id")
	if serverId == "" {
		serverId = "prisma-ca"
	}
	policy := pb.NewEmptyExtensions()
	policy.SetExtensionsRoot()
	ktArr := lib.ArrayToInterfaceArray(providers.KTArrayToKTStringArray(s.AllowedKT))
	policyData, err := pb.MakeNewExtension(ktArr)
	if err != nil {
		s.Log.Fatal("Error creating policy extension")
	}
	policy.Set("allowed_key_types", policyData)
	return &pb.ConfigReply{
		ServerProtocolVersion: &pb.Version{
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
		ReplyTime: uint64(time.Now().Unix()),
		ServerId:  serverId,
		Policy:    policy,
	}, nil
}

func (cs *CaServer) RegisterServer(s *grpc.Server) {
	pb.RegisterPrismaCaServer(s, cs)
}