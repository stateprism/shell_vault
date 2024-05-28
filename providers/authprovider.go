package providers

import (
	"context"

	pb "github.com/stateprism/prisma_ca/protocol"
)

type AuthProvider interface {
	String() string
	Authenticate(ctx context.Context, msg *pb.AuthRequest) (bool, error)
}
