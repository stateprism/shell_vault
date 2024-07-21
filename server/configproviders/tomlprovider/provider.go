package tomlprovider

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"github.com/xadaemon/libprisma/memkv"
	"github.com/xadaemon/shell_vault/server/providers"
	"os"
	"path"
)

type TomlConfigProvider struct {
	keys     *memkv.MemKV
	filename string
}

func New(filename string) (providers.ConfigurationProvider, error) {
	stat, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	data := make([]byte, stat.Size())

	fileHandle, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func(fileHandle *os.File) {
		err := fileHandle.Close()
		if err != nil {
			fmt.Println("failed to close file")
			panic(err)
		}
	}(fileHandle)

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

	store := memkv.NewMemKV(".", &memkv.Opts{CaseInsensitive: false})
	store.ImportMap(keys)

	return &TomlConfigProvider{
		keys:     store,
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
	keys, ok := p.keys.Get(key)
	if !ok {
		return nil, fmt.Errorf("key %s not found", key)
	}
	return keys, nil
}

func (p *TomlConfigProvider) GetMany(keys ...string) ([]interface{}, error) {
	ret := make([]interface{}, len(keys))

	for i, key := range keys {
		val, err := p.Get(key)
		if err != nil {
			return nil, nil
		}
		valStr := val
		ret[i] = valStr
	}
	return ret, nil
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

func (p *TomlConfigProvider) GetStrings(keys ...string) ([]string, error) {
	ret := make([]string, len(keys))

	for i, key := range keys {
		val, err := p.Get(key)
		if err != nil {
			return nil, nil
		}
		valStr, ok := val.(string)
		if !ok {
			return nil, providers.CONFIG_ERROR_INVALID_VALUE
		}
		ret[i] = valStr
	}
	return ret, nil
}

func (p *TomlConfigProvider) GetStringOrDefault(key, def string) string {
	val, err := p.Get(key)
	if err != nil {
		return def
	}
	valStr, ok := val.(string)
	if !ok {
		return def
	}
	return valStr
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
	return p.keys.Contains(key)
}

func (p *TomlConfigProvider) Set(key string, value interface{}) error {
	if !p.keys.Set(key, value) {
		return fmt.Errorf("failed to set %s", key)
	}
	return nil
}
