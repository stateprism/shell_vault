package pamprovider

import (
	"context"
	"errors"

	"github.com/msteinert/pam"
	"github.com/stateprism/prisma_ca/server/authproviders"
	pb "github.com/stateprism/prisma_ca/server/protocol"
)

type pamProvider struct{}

func New() *pamProvider {
	return &pamProvider{}
}

func (p *pamProvider) String() string {
	return "PAM"
}

func (p *pamProvider) Authenticate(ctx context.Context, msg *pb.AuthRequest) (bool, error) {
	req, err := authproviders.RequestFromBytes(msg.GetAuthRequest())
	if err != nil {
		return false, errors.New("invalid request")
	}
	password := req.Password
	username := req.Username
	t, err := pam.StartFunc("passwd", "", func(s pam.Style, _ string) (string, error) {
		switch s {
		case pam.PromptEchoOff:
			return password, nil
		case pam.PromptEchoOn:
			return username, nil
		case pam.ErrorMsg:
			return "", nil
		}
		return "", nil
	})

	if err != nil {
		return false, err
	}

	err = t.Authenticate(0)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (p *pamProvider) GetUserIdentifier(ctx context.Context, msg *pb.AuthRequest) string {
	req, _ := authproviders.RequestFromBytes(msg.GetAuthRequest())
	return req.Username
}
