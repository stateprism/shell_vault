package tomlprovider

import (
	"errors"
	"path"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/afero"
	"github.com/stateprism/prisma_ca/server/providers"
)

type TomlConfigProvider struct {
	keys     map[string]interface{}
	filename string
}

func New(fs afero.Fs, filename string) (providers.ConfigurationProvider, error) {
	stat, err := fs.Stat(filename)
	if err != nil {
		return nil, err
	}

	data := make([]byte, stat.Size())

	fileHandle, err := fs.Open(filename)
	if err != nil {
		return nil, err
	}

	defer fileHandle.Close()

	n, err := fileHandle.Read(data)
	if err != nil {
		return nil, err
	} else if n != int(stat.Size()) {
		return nil, errors.New("failed to read the entire file")
	}

	keys, err := readToml(data)
	if err != nil {
		return nil, err
	}

	return &TomlConfigProvider{
		keys:     keys,
		filename: filename,
	}, nil
}

func readToml(data []byte) (map[string]interface{}, error) {
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
	return "TOMProvider"
}

func (p *TomlConfigProvider) GetLocalStore() string {
	return path.Dir(p.filename)
}

func (p *TomlConfigProvider) Get(key string) (interface{}, error) {
	view := p.keys
	keys := strings.Split(key, ".")
	if strings.Contains(key, ".") {
		for i, k := range keys {
			if i == len(keys)-1 {
				break
			}
			viewTemp, ok := view[k].(map[string]interface{})
			if !ok {
				return "", providers.CONFIG_ERROR_INVALID_PATH
			}
			view = viewTemp
		}
	}
	key = keys[len(keys)-1]
	val, ok := view[key]
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

func (p *TomlConfigProvider) GetInt64(key string) (int64, error) {
	val, err := p.Get(key)
	if err != nil {
		return 0, nil
	}
	valInt, ok := val.(int64)
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

func (p *TomlConfigProvider) HasKey(key string) bool {
	if _, err := p.Get(key); err != nil {
		return false
	}
	return true
}

func (p *TomlConfigProvider) Set(key string, value interface{}) error {
	err := p.Set(key, value)
	if err != nil {
		return err
	}
	return nil
}
