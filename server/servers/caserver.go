package servers

import (
	"context"
	cryptorand "crypto/rand"
	"fmt"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand/v2"
	"slices"
	"time"

	"github.com/stateprism/libprisma/memkv"
	pb "github.com/stateprism/prisma_ca/rpc/caproto"
	"github.com/stateprism/prisma_ca/server/lib"
	"github.com/stateprism/prisma_ca/server/providers"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type CaServer struct {
	pb.UnsafePrismaCaServer
	Config      providers.ConfigurationProvider
	Auth        providers.AuthProvider
	DB          providers.DatabaseProvider
	KProvider   providers.KeychainProvider
	AllowedKT   []providers.KeyType
	Env         *providers.EnvProvider
	Log         *zap.Logger
	Listen      string
	MemKV       *memkv.MemKV
	rootTtl     int64
	rootKt      providers.KeyType
	userCertTtl uint64
}

type CaServerParams struct {
	fx.In
	Logger    *zap.Logger
	Config    providers.ConfigurationProvider
	Auth      providers.AuthProvider
	Env       *providers.EnvProvider
	KProvider providers.KeychainProvider
	MemKV     *memkv.MemKV
}

func NewCAServer(p CaServerParams) (*CaServer, error) {
	s := &CaServer{
		Config:    p.Config,
		Auth:      p.Auth,
		Log:       p.Logger,
		KProvider: p.KProvider,
		MemKV:     p.MemKV,
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
	// TODO: Handle these possible errors
	rKts, _ := s.Config.GetString("ca_server.root_key_type")
	rktTtl, _ := s.Config.GetInt64("ca_server.root_key_max_ttl")
	userCertTtl, _ := s.Config.GetInt64("ca_server.user_cert_ttl")
	s.rootKt = providers.KTFromString(rKts)
	s.rootTtl = rktTtl
	s.userCertTtl = uint64(userCertTtl)

	s.KProvider.SetExpKeyHook(s.expKeyHook)

	if err := s.rotateRootKey(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *CaServer) Authenticate(ctx context.Context, _ *pb.EmptyMsg) (*pb.AuthReply, error) {
	ent, err := s.Auth.Authenticate(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.AuthReply{
		AuthTime:  time.Now().Unix(),
		AuthUntil: time.Now().Add(time.Duration(s.userCertTtl) * time.Second).Unix(),
		AuthToken: ent,
		Success:   true,
		Errors:    nil,
	}, nil
}

func (s *CaServer) RequestCert(ctx context.Context, msg *pb.CertRequest) (*pb.CertReply, error) {
	pko, err := s.KProvider.RetrieveKey("rootKey")
	if err != nil {
		s.Log.Fatal("Root key not found", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	pkr := pko.UnwrapKey()
	signer, err := ssh.NewSignerFromKey(pkr)
	if err != nil {
		s.Log.Fatal("Error creating signer", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	session := ctx.Value("session").(*memkv.MemKV)
	if session == nil {
		return nil, fmt.Errorf("session is nil")
	}

	pubk, err := ssh.ParsePublicKey(msg.PublicKey)
	if err != nil {
		return nil, err
	}
	serial := rand.Uint64()
	principal, ok := session.Get("session.user.username")
	if !ok {
		return nil, fmt.Errorf("principal not found")
	}
	cert := ssh.Certificate{
		Key:         pubk,
		Serial:      serial,
		ValidAfter:  uint64(time.Now().Add(1 * time.Second).Unix()),
		ValidBefore: uint64(time.Now().Add(time.Duration(s.userCertTtl) * time.Second).Unix()),
		CertType:    ssh.UserCert,
		ValidPrincipals: []string{
			principal.(string),
		},
	}
	err = cert.SignCert(cryptorand.Reader, signer)
	if err != nil {
		s.Log.Error("Error signing cert", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal error")
	}
	return &pb.CertReply{
		Cert:       string(ssh.MarshalAuthorizedKey(&cert)),
		ValidUntil: cert.ValidBefore,
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

func (s *CaServer) RegisterServer(srv *grpc.Server) {
	pb.RegisterPrismaCaServer(srv, s)
}

func (s *CaServer) rotateRootKey() error {
	key, err := s.KProvider.RetrieveKey("rootKey")
	if err != nil && err.Error() != "sql: no rows in result set" {
		return err
	}
	if (err != nil && err.Error() == "sql: no rows in result set") || time.Now().Unix() > time.Now().Add(key.GetTtl()).Unix() {
		s.Log.Info("Rotating root key")
		_, err := s.KProvider.MakeAndReplaceKey("rootKey", s.rootKt, s.rootTtl)
		if err != nil {
			return err
		}
	}
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
