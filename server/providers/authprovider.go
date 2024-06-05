package providers

import (
	"context"
	"github.com/stateprism/libprisma/memkv"
)

type EntityInfo struct {
	*memkv.MemKV
}

type AuthProvider interface {
	String() string
	Authenticate(ctx context.Context) (string, error)
	GetSession(ctx context.Context) (context.Context, error)
}
