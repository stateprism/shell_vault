package providers

import "errors"

type ConfigError int

const (
	CONFIG_ERROR_UNKNOWN        ConfigError = iota
	CONFIG_ERROR_INVALID_SOURCE ConfigError = iota
	CONFIG_ERROR_INVALID_KEY    ConfigError = iota
	CONFIG_ERROR_INVALID_VALUE  ConfigError = iota
	CONFIG_ERROR_INVALID_PATH   ConfigError = iota
)

func (e ConfigError) Error() string {
	switch e {
	case CONFIG_ERROR_INVALID_SOURCE:
		return "Invalid configuration source"
	case CONFIG_ERROR_INVALID_KEY:
		return "Invalid configuration key"
	case CONFIG_ERROR_INVALID_VALUE:
		return "Invalid configuration value"
	case CONFIG_ERROR_INVALID_PATH:
		return "Invalid configuration path"
	default:
		return "Unknown error"
	}
}

func GetManyFromProvider[T any](p ConfigurationProvider, keys ...string) ([]T, error) {
	if p == nil {
		return nil, errors.New("provider is nil")
	}

	ret := make([]T, len(keys))

	for i, key := range keys {
		val, err := p.Get(key)
		if err != nil {
			return nil, nil
		}
		valStr, ok := val.(T)
		if !ok {
			return nil, CONFIG_ERROR_INVALID_VALUE
		}
		ret[i] = valStr
	}
	return ret, nil
}

func GetOrDefault[T any](p ConfigurationProvider, key string, def T) T {
	if p == nil {
		panic("provider is nil")
	}

	v, err := p.Get(key)
	v, ok := v.(T)
	if err != nil || !ok {
		return def
	}

	return v.(T)
}

type ConfigurationProvider interface {
	String() string
	IsWriteable() bool
	GetLocalStore() string
	HasKey(key string) bool
	GetString(key string) (string, error)
	GetStrings(keys ...string) ([]string, error)
	GetStringOrDefault(key, def string) string
	GetInt(key string) (int, error)
	GetInt64(key string) (int64, error)
	GetBytes(key string) ([]byte, error)
	GetBool(key string) (bool, error)
	Get(key string) (interface{}, error)
	GetMany(keys ...string) ([]interface{}, error)
	Set(key string, val interface{}) error
}
