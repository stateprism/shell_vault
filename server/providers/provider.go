package providers

import (
	"fmt"
	"strings"

	"github.com/spf13/afero"
	"github.com/stateprism/prisma_ca/server/authproviders/pamprovider"
	"github.com/stateprism/prisma_ca/server/authproviders/plainfileprovider"
)

func NewAuthProvider(config ConfigurationProvider) (AuthProvider, error) {
	providerName, err := config.GetString("ca_server.auth_provider")
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(providerName, "builtin://") {
		striped := strings.TrimPrefix(providerName, "builtin://")
		split := strings.Split(striped, "?")
		switch split[0] {
		case "pam":
			return pamprovider.New(), nil
		case "plainfile":
			fs := afero.NewOsFs()
			pro, err := plainfileprovider.New(fs, split[1])
			if err != nil {
				return nil, err
			}
			return pro, nil
		default:
			return nil, fmt.Errorf("unknown builtin auth provider")
		}
	}
	return nil, fmt.Errorf("unknown auth provider")
}
