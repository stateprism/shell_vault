package providers

import (
	"context"

	pb "github.com/stateprism/prisma_ca/rpc/caproto"
)

type AuthProvider interface {
	String() string
	Authenticate(ctx context.Context, msg *pb.AuthRequest) (bool, error)
	GetUserIdentifier(ctx context.Context, msg *pb.AuthRequest) string
}
