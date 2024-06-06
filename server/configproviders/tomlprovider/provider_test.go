package tomlprovider_test

import (
	"testing"

	_ "embed"

	"github.com/spf13/afero"
	"github.com/stateprism/shell_vault/server/configproviders/tomlprovider"
	"github.com/stateprism/shell_vault/server/providers"
)

//go:embed config.toml
var config []byte

func TestJsonProviderGet(t *testing.T) {
	tests := []struct {
		Name         string
		FileContents []byte
		Key          string
		Expect       interface{}
		Err          error
	}{
		{
			Name:         "Success",
			FileContents: config,
			Key:          "ca_server.ca_host",
			Expect:       "localhost:5000",
			Err:          nil,
		},
		{
			Name:         "Fetch non-existent key",
			FileContents: config,
			Key:          "SomeKeys",
			Expect:       "",
			Err:          providers.CONFIG_ERROR_INVALID_KEY,
		},
	}

	for _, tt := range tests {
		fs := afero.NewMemMapFs()
		file, _ := fs.Create("config.toml")
		file.Write(tt.FileContents)

		provider, err := tomlprovider.New(fs, file.Name())
		if err != nil {
			t.Fatalf("error creating the jsonprovider, but got: %s", err)
		}
		val, err := provider.Get(tt.Key)
		if tt.Err == nil && err != tt.Err {
			t.Fatalf("unexpected error, `%s`", err)
		} else if err != tt.Err {
			t.Fatalf("expected error `%s`, but got `%s`", tt.Err, err)
		}
		if tt.Expect != val {
			t.Fatalf("returned value `%s` doesn't match the expected `%s`", val, tt.Expect)
		}
	}
}
