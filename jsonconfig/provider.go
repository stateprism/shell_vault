package jsonprovider

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/stateprism/prism_ca/config"
)

type JsonConfigProvider struct {
	keys     map[string]interface{}
	filename string
}

func New(filename string) (*JsonConfigProvider, error) {
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

	return &JsonConfigProvider{
		keys:     keys,
		filename: filename,
	}, nil
}

func ReadJson(data []byte) (map[string]interface{}, error) {
	var decoded map[string]interface{}
	err := json.Unmarshal(data, &decoded)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func (p *JsonConfigProvider) IsWriteable() bool {
	return false
}

func (p *JsonConfigProvider) String() string {
	return "JsonConfigProvider"
}

func (p *JsonConfigProvider) Get(key string) (interface{}, error) {
	val, ok := p.keys[key]
	if !ok {
		return "", config.NewConfigError(config.CONFIG_ERROR_INVALID_KEY, p)
	}
	return val, nil
}

func (p *JsonConfigProvider) GetString(key string) (string, error) {
	val, err := p.Get(key)
	if err != nil {
		return "", nil
	}
	valStr, ok := val.(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("%s value is not a string", key))
	}
	return valStr, nil
}

func (p *JsonConfigProvider) GetInt(key string) (int, error) {
	val, err := p.Get(key)
	if err != nil {
		return 0, nil
	}
	valInt, ok := val.(int)
	if !ok {
		return 0, errors.New(fmt.Sprintf("%s value is not a string", key))
	}
	return valInt, nil
}

func (p *JsonConfigProvider) Set(key string, value interface{}) error {
	return errors.New("config provider JsonProvider is not writeable")
}
