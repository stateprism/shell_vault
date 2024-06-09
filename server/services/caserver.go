package services

import (
	"context"
	cryptorand "crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"github.com/stateprism/shell_vault/rpc/common"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand/v2"
	"slices"
	"time"

	pb "github.com/stateprism/shell_vault/rpc/caproto"
	pbcommon "github.com/stateprism/shell_vault/rpc/common"
	"github.com/stateprism/shell_vault/server/lib"
	"github.com/stateprism/shell_vault/server/providers"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type CaServer struct {
	pb.CertificateAuthorityServer
	config      providers.ConfigurationProvider
	authPro     providers.AuthProvider
	vault       providers.KeychainProvider
	allowedKT   []providers.KeyType
	logger      *zap.Logger
	Listen      string
	rootTtl     int64
	rootKt      providers.KeyType
	userCertTtl uint64
	sessionTtl  time.Duration
}

type CaServerParams struct {
	fx.In
	LC        fx.Lifecycle
	Logger    *zap.Logger
	Config    providers.ConfigurationProvider
	Auth      providers.AuthProvider
	Env       *providers.EnvProvider
	KProvider providers.KeychainProvider
}

func NewCAServer(p CaServerParams) (*CaServer, error) {
	s := &CaServer{
		config:  p.Config,
		authPro: p.Auth,
		logger:  p.Logger,
		vault:   p.KProvider,
	}
	if s.config == nil {
		return nil, fmt.Errorf("config is nil")
	}
	akt, err := s.config.Get("ca_server.allowed_key_types")
	if err != nil {
		return nil, err
	}
	kt, _ := akt.([]interface{})
	s.allowedKT, _, _ = providers.KTStringArrayToKTArray(lib.InterfaceArrayToArray[string](kt))
	s.Listen, err = s.config.GetString("ca_server.listen")
	if err != nil {
		return nil, err
	}

	if s.authPro == nil {
		return nil, fmt.Errorf("auth is nil")
	}
	// TODO: Handle these possible errors
	rKts, _ := s.config.GetString("ca_server.root_key_type")
	rktTtl, _ := s.config.GetInt64("ca_server.root_key_max_ttl")
	userCertTtl, _ := s.config.GetInt64("ca_server.user_cert_ttl")
	s.rootKt = providers.KTFromString(rKts)
	s.rootTtl = rktTtl
	s.userCertTtl = uint64(userCertTtl)
	sessionTtl := providers.GetOrDefault(s.config, "ca_server.auth_session_ttl", int64(72000))
	s.sessionTtl = time.Duration(sessionTtl)

	s.vault.SetExpKeyHook(s.expKeyHook)

	p.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			s.logger.Info("Checking for root key rotation necessity on startup")
			if err := s.rotateRootKey(); err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})

	return s, nil
}

func (s *CaServer) Authenticate(ctx context.Context, _ *pbcommon.Empty) (*pbcommon.AuthReply, error) {
	ctx = context.WithValue(ctx, "sessionTtl", s.sessionTtl)
	ent, err := s.authPro.Authenticate(ctx)
	if err != nil {
		return nil, err
	}
	return &pbcommon.AuthReply{
		AuthTime:  time.Now().Unix(),
		AuthUntil: time.Now().Add(s.sessionTtl * time.Second).Unix(),
		AuthToken: ent,
		Success:   true,
	}, nil
}

func (s *CaServer) GetCurrentKey(ctx context.Context, _ *pbcommon.Empty) (*pb.CertReply, error) {
	pko, err := s.vault.RetrieveKey("rootKey")
	if err != nil {
		s.logger.Fatal("Root key not found", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	pk, err := ssh.NewPublicKey(pko.GetPublicKey())
	if err != nil {
		s.logger.Fatal("Error getting public key", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}

	ttl := time.Now().Add(pko.GetTtl() * time.Second).UTC().Unix()

	return &pb.CertReply{
		Cert:       string(ssh.MarshalAuthorizedKey(pk)),
		ValidUntil: ttl,
	}, nil
}

func (s *CaServer) RequestUserCertificate(ctx context.Context, msg *pb.UserCertRequest) (*pb.CertReply, error) {
	// Retrieve the key
	pko, err := s.vault.RetrieveKey("rootKey")
	if err != nil {
		s.logger.Fatal("Root key not found", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	// Get the key into a type we can use with ssh lib
	pkr := pko.UnwrapKey()
	signer, err := ssh.NewSignerFromKey(pkr)
	if err != nil {
		s.logger.Fatal("Error creating signer", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	session := ctx.Value("session").(*providers.SessionInfo)
	if session == nil {
		return nil, fmt.Errorf("session is nil")
	}
	// Load the key into the signer
	pubk, err := ssh.ParsePublicKey(msg.PublicKey)
	if err != nil {
		return nil, err
	}
	// Make the certificate
	serial := rand.Uint64()
	cert := ssh.Certificate{
		Key:         pubk,
		Serial:      serial,
		ValidAfter:  uint64(time.Now().Add(1 * time.Second).UTC().Unix()),
		ValidBefore: uint64(time.Now().Add(time.Duration(s.userCertTtl) * time.Second).UTC().Unix()),
		CertType:    ssh.UserCert,
		ValidPrincipals: []string{
			session.Principal,
			uuid.New().URN(),
		},
	}
	err = cert.SignCert(cryptorand.Reader, signer)
	if err != nil {
		s.logger.Error("Error signing cert", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &pb.CertReply{
		Cert:       string(ssh.MarshalAuthorizedKey(&cert)),
		ValidUntil: int64(cert.ValidBefore),
	}, nil
}

func (s *CaServer) RequestServerCertificate(ctx context.Context, msg *pb.HostCertRequest) (*pb.CertReply, error) {
	// Retrieve the key
	pko, err := s.vault.RetrieveKey("rootKey")
	if err != nil {
		s.logger.Fatal("Root key not found", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	// Get the key into a type we can use with ssh lib
	pkr := pko.UnwrapKey()
	signer, err := ssh.NewSignerFromKey(pkr)
	if err != nil {
		s.logger.Fatal("Error creating signer", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	session := ctx.Value("session").(*providers.SessionInfo)
	if session == nil {
		return nil, fmt.Errorf("session is nil")
	}
	// Load the key into the signer
	pubk, err := ssh.ParsePublicKey(msg.PublicKey)
	if err != nil {
		return nil, err
	}
	serial := rand.Uint64()
	cert := ssh.Certificate{
		Key:             pubk,
		Serial:          serial,
		ValidAfter:      uint64(time.Now().Add(1 * time.Second).UTC().Unix()),
		ValidBefore:     uint64(time.Now().Add(time.Duration(s.userCertTtl) * time.Second).UTC().Unix()),
		CertType:        ssh.HostCert,
		ValidPrincipals: msg.GetHostnames(),
	}
	err = cert.SignCert(cryptorand.Reader, signer)
	if err != nil {
		s.logger.Error("Error signing cert", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &pb.CertReply{
		Cert:       string(ssh.MarshalAuthorizedKey(&cert)),
		ValidUntil: int64(cert.ValidBefore),
	}, nil
}

func (s *CaServer) GetConfig(context.Context, *pb.ConfigRequest) (*pb.ConfigReply, error) {
	serverId, _ := s.config.GetString("server_id")
	if serverId == "" {
		serverId = "prisma-ca"
	}
	policy := common.NewEmptyExtensions()
	policy.SetExtensionsRoot()
	ktArr := lib.ArrayToInterfaceArray(providers.KTArrayToKTStringArray(s.allowedKT))
	policyData, err := common.MakeNewExtension(ktArr)
	if err != nil {
		s.logger.Fatal("Error creating policy extension")
	}
	policy.Set("allowed_key_types", policyData)
	return &pb.ConfigReply{
		ServerProtocolVersion: &pbcommon.Version{
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
		ReplyTime: uint64(time.Now().Unix()),
		ServerId:  serverId,
		Policy:    policy,
	}, nil
}

func (s *CaServer) RegisterServer(srv *grpc.Server) {
	pb.RegisterCertificateAuthorityServer(srv, s)
}

func (s *CaServer) rotateRootKey() error {
	key, err := s.vault.RetrieveKey("rootKey")
	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}
	if err != nil && err.Error() == "sql: no rows in result set" {
		s.logger.Info("Root key not found, creating new key")
		_, err := s.vault.MakeAndReplaceKey("rootKey", s.rootKt, s.rootTtl)
		if err != nil {
			return err
		}
		return nil
	}
	if time.Now().Unix() > time.Now().Add(key.GetTtl()).Unix() {
		s.logger.Info("Rotating root key")
		_, err := s.vault.MakeAndReplaceKey("rootKey", s.rootKt, s.rootTtl)
		if err != nil {
			return err
		}
		return nil
	}
	s.logger.Info("Root key does not need to be rotated")
	return nil
}

func (s *CaServer) expKeyHook(p providers.KeychainProvider, ids []providers.KeyIdentifier) {
	if slices.Contains(ids, "rootKey") {
		err := s.rotateRootKey()
		if err != nil {
			panic("Failed to rotate root keys")
		}
	}
}
