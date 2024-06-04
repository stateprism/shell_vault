package pamprovider

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/stateprism/libprisma/memkv"
	"github.com/stateprism/prisma_ca/server/middleware"
	"github.com/stateprism/prisma_ca/server/providers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/msteinert/pam"
)

type PamProvider struct {
	conf providers.ConfigurationProvider
	kv   *memkv.MemKV
}

func New(config providers.ConfigurationProvider, kv *memkv.MemKV) providers.AuthProvider {
	return &PamProvider{
		conf: config,
		kv:   kv,
	}
}

func (p *PamProvider) String() string {
	return "PAM"
}

func (p *PamProvider) Authenticate(ctx context.Context, authString string) (*providers.EntityInfo, error) {
	token, err := middleware.AuthFromMetadata(ctx, "pam", "authorization")
	if err != nil {
		return nil, err
	}

	tokenBytes := make([]byte, 0)
	_, err = base64.URLEncoding.Decode(tokenBytes, []byte(token))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Your request is not acceptable by any of the enabled auth providers")
	}
	username, password, found := bytes.Cut(tokenBytes, []byte{0x1E})
	if !found {
		return nil, status.Error(codes.Unauthenticated, "Your request is not acceptable by any of the enabled auth providers")
	}
	t, err := pam.StartFunc("passwd", "", func(s pam.Style, _ string) (string, error) {
		switch s {
		case pam.PromptEchoOff:
			return string(password), nil
		case pam.PromptEchoOn:
			return string(username), nil
		case pam.ErrorMsg:
			return "", nil
		}
		return "", nil
	})
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "The supplied credentials are invalid")
	}

	err = t.Authenticate(0)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "The supplied credentials are invalid")
	}

	sesId := uuid.New()
	entInfo := memkv.NewMemKV(".", nil)
	entInfo.Set("user.username", username)
	entInfo.Set("user.realm", "linuxPAM")
	entInfo.Set("session.id", sesId)

	p.kv.Set("sessions."+sesId.String(), entInfo)

	return &providers.EntityInfo{MemKV: entInfo}, nil
}
