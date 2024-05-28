package config

import "fmt"

type ConfigErrorKind int

const (
	CONFIG_ERROR_INVALID_SOURCE ConfigErrorKind = 0
	CONFIG_ERROR_INVALID_KEY    ConfigErrorKind = 1
	CONFIG_ERROR_INVALID_VALUE  ConfigErrorKind = 2
)

type ConfigError struct {
	Kind     ConfigErrorKind
	Provider ConfigurationProvider
}

func NewConfigError(kind ConfigErrorKind, provider ConfigurationProvider) *ConfigError {
	return &ConfigError{
		Kind:     kind,
		Provider: provider,
	}
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("provider %s returned code %d", e.Provider.String(), e.Kind)
}

type ConfigurationProvider interface {
	String() string
	IsWriteable() bool
	GetString(key string) (string, error)
	GetInt(key string) (int, error)
	Get(key string) (interface{}, error)
	Set(key string, val interface{}) error
}
