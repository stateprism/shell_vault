package providers

import (
	"os"
	"strings"
)

type EnvProvider struct {
	// filter for environment variables
	prefix string
	// copy of the os.Environ() slice made at initialization, filtered by prefix
	env map[string]string
}

// NewEnvProvider creates a new EnvProvider with the given prefix
func NewEnvProvider(prefix string) *EnvProvider {
	envCopy := make(map[string]string)
	for _, env := range os.Environ() {
		split := strings.Split(env, "=")
		if strings.HasPrefix(split[0], prefix) {
			envCopy[split[0]] = split[1]
		}
	}

	return &EnvProvider{
		prefix: prefix,
		env:    envCopy,
	}
}

// GetEnv returns the value of the environment variable with the given key, and a boolean indicating if the key was found
func (p *EnvProvider) GetEnv(key string) (*string, bool) {
	val, ok := p.env[p.prefix+key]
	return &val, ok
}

func (p *EnvProvider) GetEnvOrDefault(key, def string) string {
	val, ok := p.env[p.prefix+key]
	if !ok {
		return def
	}
	return val
}

func (p *EnvProvider) IsEnvEqual(key, value string) bool {
	val, ok := p.env[p.prefix+key]
	if !ok {
		return false
	}
	return val == value
}

func (p *EnvProvider) SetEnv(key string, value string) {
	err := os.Setenv(p.prefix+key, value)
	if err != nil {
		panic(err)
	}
}

func (p *EnvProvider) UnsetEnv(key string) {
	err := os.Unsetenv(p.prefix + key)
	if err != nil {
		panic(err)
	}
}

func (p *EnvProvider) GetEnvMap() map[string]string {
	return p.env
}
