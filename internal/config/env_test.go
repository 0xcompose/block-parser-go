package config

import (
	"os"
	"testing"
)

func TestGetSafeEnvValue(t *testing.T) {
	os.Setenv("TEST_ENV", "test")

	value := GetEnvValueSafely("TEST_ENV", "TEST_ENV is not set")

	if value != "test" {
		t.Errorf("Expected value to be 'test', but got '%s'", value)
	}

	t.Cleanup(os.Clearenv)
}
