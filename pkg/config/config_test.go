package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("loads a testing value from the .env file", func(t *testing.T) {
		LoadConfigTest()

		got := os.Getenv("TEST_VALUE")
		want := "xDLUL420"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}
