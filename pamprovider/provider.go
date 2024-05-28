package pamprovider

import (
	"context"

	"github.com/msteinert/pam"
	pb "github.com/stateprism/prisma_ca/protocol"
	"github.com/stateprism/prisma_ca/providers"
)

type pamProvider struct {
	Logger providers.LogProvider
}

func New(logger providers.LogProvider) *pamProvider {
	return &pamProvider{
		Logger: logger,
	}
}

func (p *pamProvider) String() string {
	return "PAM"
}

func (p *pamProvider) Authenticate(ctx context.Context, msg *pb.AuthRequest) (bool, error) {
	if msg.Password == nil || msg.Username == nil {
		p.Logger.Log(providers.LOG_LEVEL_INFO, "Username or password not provided")
		return false, nil
	}
	password := *msg.Password
	username := *msg.Username
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
