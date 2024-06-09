package services

import (
	"context"
	pbcommon "github.com/stateprism/shell_vault/rpc/common"
	"time"

	"github.com/google/uuid"
	pb "github.com/stateprism/shell_vault/rpc/adminproto"
	"github.com/stateprism/shell_vault/server/providers"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AdminServerParams struct {
	fx.In
	Config  providers.ConfigurationProvider
	AuthPro providers.AuthProvider
	Env     *providers.EnvProvider
	Logger  *zap.Logger
}

type AuthChallenge struct {
	Id        uuid.UUID
	Challenge []byte
	Expires   int64
	Nonce     uint64
}

type AdminServer struct {
	pb.UnsafeAdminServiceServer
	config     providers.ConfigurationProvider
	authPro    providers.AuthProvider
	logger     *zap.Logger
	Listen     string
	sessionTtl time.Duration
}

func NewAdminServer(p AdminServerParams) (*AdminServer, error) {
	s := new(AdminServer)
	s.config = p.Config
	s.logger = p.Logger
	s.authPro = p.AuthPro

	s.Listen = providers.GetOrDefault(s.config, "auth_server.listen", "localhost:8000")
	sessionTtl := providers.GetOrDefault(s.config, "ca_server.auth_session_ttl", int64(72000))
	s.sessionTtl = time.Duration(sessionTtl)

	return s, nil
}

func (a *AdminServer) Authenticate(ctx context.Context, _ *pbcommon.Empty) (*pbcommon.AuthReply, error) {
	ctx = context.WithValue(ctx, "sessionTtl", a.sessionTtl)
	ent, err := a.authPro.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	return &pbcommon.AuthReply{
		AuthTime:  time.Now().Unix(),
		AuthUntil: time.Now().Add(a.sessionTtl * time.Second).Unix(),
		AuthToken: ent,
		Success:   true,
	}, nil
}

func (a *AdminServer) RestartServer(ctx context.Context, request *pb.StopServerRequest) (*pbcommon.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdminServer) ChangeRootCert(context.Context, *pb.ChangeRootCertRequest) (*pb.ChangeRootCertResponse, error) {
	panic("unimplemented")
}

func (a *AdminServer) AddUser(ctx context.Context, r *pb.AddUserRequest) (*pb.UserActionResponse, error) {
	return nil, nil
}

func (a *AdminServer) RegisterServer(s *grpc.Server) {
	pb.RegisterAdminServiceServer(s, a)
}
