package providers

import (
	"context"
	"github.com/google/uuid"
)

type SessionInfo struct {
	Principal string
	ValidTo   int64
	Realm     string
	Id        uuid.UUID
	Deadline  int64
	ExtraData map[string]any
}

type AuthProvider interface {
	String() string
	Authenticate(ctx context.Context) (string, error)
	Authorize(ctx context.Context, method string) (context.Context, error)
}
