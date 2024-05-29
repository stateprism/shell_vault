package tomlprovider_test

import (
	"os"
	"testing"

	"github.com/stateprism/prisma_ca/providers"
	"github.com/stateprism/prisma_ca/tomlprovider"
)

func TestJsonProviderGet(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		Name     string
		JsonData []byte
		Key      string
		Expect   interface{}
		Err      error
	}{
		{
			Name:     "Success",
			JsonData: []byte(`{"SomeKey": "test string"}`),
			Key:      "SomeKey",
			Expect:   "test string",
			Err:      nil,
		},
		{
			Name:     "Fetch non-existent key",
			JsonData: []byte(`{"SomeKey": "test string"}`),
			Key:      "SomeKeys",
			Expect:   "",
			Err:      providers.CONFIG_ERROR_INVALID_KEY,
		},
	}

	for _, tt := range tests {
		file, err := os.CreateTemp(tmpDir, "Test_")

		if err != nil {
			t.Fatal("failed to create the tempfile")
		}
		file.Write(tt.JsonData)
		provider, err := tomlprovider.New(file.Name())
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

		os.Remove(file.Name())
	}
}
