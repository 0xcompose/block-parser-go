package config

import (
	"os"
	"testing"
)

// TestModuleInitLoadsEnv tests that the `env` module from `config` package loads the environment variables from the .env file
func TestModuleInitLoadsEnv(t *testing.T) {
	GetEnv("ALCHEMY_API_KEY", "ALCHEMY_API_KEY is not set")

	if len(os.Getenv("ALCHEMY_API_KEY")) == 0 {
		t.Errorf("Expected env module to load environment variables, but received empty ALCHEMY_API_KEY")
	}
}

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_ENV", "test")

	value, err := GetEnv("TEST_ENV", "TEST_ENV is not set")

	if err != nil {
		t.Errorf("Expected error to be nil, but got '%v'", err)
	}

	if value != "test" {
		t.Errorf("Expected value to be 'test', but got '%s'", value)
	}

	t.Cleanup(os.Clearenv)
}

func TestGetEnvWithDefault(t *testing.T) {
	value := GetEnvWithDefault("SOME_RANDOM_ENV", "test")

	if value != "test" {
		t.Errorf("Expected value to be 'test', but got '%s'", value)
	}

	os.Setenv("SOME_RANDOM_ENV", "some_value")
	value = GetEnvWithDefault("SOME_RANDOM_ENV", "test")

	if value != "some_value" {
		t.Errorf("Expected value to be 'some_value', but got '%s'", value)
	}

	t.Cleanup(os.Clearenv)
}
