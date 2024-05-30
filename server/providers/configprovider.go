package providers

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

type ConfigurationProvider interface {
	String() string
	IsWriteable() bool
	GetString(key string) (string, error)
	GetInt(key string) (int, error)
	GetBytes(key string) ([]byte, error)
	GetBool(key string) (bool, error)
	Get(key string) (interface{}, error)
	Set(key string, val interface{}) error
}
