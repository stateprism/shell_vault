package server

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"time"

	"github.com/stateprism/prisma_ca/pamprovider"
	pb "github.com/stateprism/prisma_ca/protocol"
	"github.com/stateprism/prisma_ca/providers"
)

type caServer struct {
	pb.UnsafePrismaCaServer
	Logger  providers.LogProvider
	Config  providers.ConfigurationProvider
	Auth    providers.AuthProvider
	DB      providers.DatabaseProvider
	Env     *providers.EnvProvider
	HmacKey []byte
}

func NewCAServer(configProvider providers.ConfigurationProvider, envProvider *providers.EnvProvider, log providers.LogProvider) pb.PrismaCaServer {
	server := new(caServer)
	server.Config = configProvider
	server.Logger = log
	server.Env = envProvider
	key, err := configProvider.GetString("hmac_key")

	if err != nil {
		server.Logger.Fatalf("Error reading HMAC key: %v", err)
	}

	authProvider, err := configProvider.GetString("auth_provider")
	if err != nil {
		server.Logger.Fatalf("Error reading auth provider: %v", err)
	}
	switch authProvider {
	case "pam":
		server.Auth = pamprovider.New(server.Logger)
	default:
		server.Logger.Fatalf("Unknown auth provider: %s", authProvider)
	}

	server.Logger.Logf(providers.LOG_LEVEL_INFO, "Using auth provider: %s", server.Auth.String())

	server.HmacKey = []byte(key)
	return server
}

func (s *caServer) Authenticate(ctx context.Context, msg *pb.AuthRequest) (*pb.AuthReply, error) {
	authTime := uint64(time.Now().Unix())
	authUntil := uint64(time.Now().Add(time.Hour).Unix())
	authSuccess, err := s.Auth.Authenticate(ctx, msg)
	if err != nil {
		return &pb.AuthReply{
			Success:  false,
			AuthTime: authTime,
			Errors: &pb.Errors{
				Errors: map[string]string{"AuthError": err.Error()},
			},
		}, nil
	}
	if authSuccess {
		token := generateToken(msg.Username, s.HmacKey)
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

func generateToken(username *string, b []byte) string {
	var token []byte

	time := time.Now().Unix()
	timeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timeBytes, uint64(time))
	token = append(token, timeBytes...)
	token = append(token, []byte(*username)...)
	hmac := hmac.New(sha256.New, b)
	hmac.Write(token)
	bytes := hmac.Sum(nil)
	return hex.EncodeToString(bytes)
}

func (s *caServer) RequestCert(ctx context.Context, msg *pb.CertRequest) (*pb.CertReply, error) {
	return &pb.CertReply{
		Cert:       []byte("cert"),
		ValidUntil: uint64(time.Now().Add(time.Hour).Unix()),
	}, nil
}

func (s *caServer) GetConfig(context.Context, *pb.ConfigRequest) (*pb.ConfigReply, error) {
	serverId, err := s.Config.GetString("server_id")
	if err != nil {
		return nil, err
	}
	return &pb.ConfigReply{
		ServerProtocolVersion: &pb.Version{
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
		ReplyTime: uint64(time.Now().Unix()),
		ServerId:  serverId,
	}, nil
}
