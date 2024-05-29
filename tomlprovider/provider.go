package tomlprovider

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/stateprism/prisma_ca/providers"
)

type TomlConfigProvider struct {
	keys     map[string]interface{}
	filename string
}

func New(filename string) (*TomlConfigProvider, error) {
	if _, err := os.Stat(filename); err != nil {
		return nil, err
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	keys, err := ReadJson(file)
	if err != nil {
		return nil, err
	}

	return &TomlConfigProvider{
		keys:     keys,
		filename: filename,
	}, nil
}

func ReadJson(data []byte) (map[string]interface{}, error) {
	var decoded map[string]interface{}
	err := toml.Unmarshal(data, &decoded)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func (p *TomlConfigProvider) IsWriteable() bool {
	return false
}

func (p *TomlConfigProvider) String() string {
	return "JsonConfigProvider"
}

func (p *TomlConfigProvider) Get(key string) (interface{}, error) {
	val, ok := p.keys[key]
	if !ok {
		return "", providers.CONFIG_ERROR_INVALID_KEY
	}
	return val, nil
}

func (p *TomlConfigProvider) GetString(key string) (string, error) {
	val, err := p.Get(key)
	if err != nil {
		return "", nil
	}
	valStr, ok := val.(string)
	if !ok {
		return "", providers.CONFIG_ERROR_INVALID_VALUE
	}
	return valStr, nil
}

func (p *TomlConfigProvider) GetInt(key string) (int, error) {
	val, err := p.Get(key)
	if err != nil {
		return 0, nil
	}
	valInt, ok := val.(int)
	if !ok {
		return 0, providers.CONFIG_ERROR_INVALID_VALUE
	}
	return valInt, nil
}

func (p *TomlConfigProvider) GetBytes(key string) ([]byte, error) {
	val, err := p.Get(key)
	if err != nil {
		return nil, nil
	}
	valBytes, ok := val.([]byte)
	if !ok {
		return nil, providers.CONFIG_ERROR_INVALID_VALUE
	}
	return valBytes, nil
}

func (p *TomlConfigProvider) GetBool(key string) (bool, error) {
	val, err := p.Get(key)
	if err != nil {
		return false, nil
	}
	valBool, ok := val.(bool)
	if !ok {
		return false, providers.CONFIG_ERROR_INVALID_VALUE
	}
	return valBool, nil
}

func (p *TomlConfigProvider) Set(key string, value interface{}) error {
	return errors.New("config provider JsonProvider is not writeable")
}
