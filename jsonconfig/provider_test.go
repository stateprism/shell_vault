package jsonprovider

import (
	"os"
	"testing"

	"github.com/stateprism/prism_ca/config"
)

var staticProvider *JsonConfigProvider = &JsonConfigProvider{}

func TestJsonProvider(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		JsonData []byte
		Key      string
		Expect   interface{}
		Err      error
	}{
		{
			JsonData: []byte(`{"SomeKey": "test string"}`),
			Key:      "SomeKey",
			Expect:   "test string",
			Err:      nil,
		},
		{
			JsonData: []byte(`{"SomeKey": "test string"}`),
			Key:      "SomeKeys",
			Expect:   "",
			Err:      config.NewConfigError(config.CONFIG_ERROR_INVALID_KEY, staticProvider),
		},
	}

	for _, tt := range tests {
		file, err := os.CreateTemp(tmpDir, "Test_")
		if err != nil {
			t.Fatal("failed to create the tempfile")
		}
		file.Write(tt.JsonData)
		provider, err := New(file.Name())
		if err != nil && tt.Err == nil {
			t.Fatalf("test didn't expect error, but got: %s", err)
		}
		val, err := provider.Get(tt.Key)
		if tt.Expect != val {
			t.Fatalf("returned value `%s` doesn't match the expected `%s`", val, tt.Expect)
		}

		os.Remove(file.Name())
	}
}
