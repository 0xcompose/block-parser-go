package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const envFilePath = "../../.env"

// as soon as the package is imported, the .env file is loaded
func init() {
	err := godotenv.Load(envFilePath)
	log.Printf("Loading .env file: %q", envFilePath)

	if err != nil {
		log.Printf("Error loading .env file: %q", err)
	}

	log.Print("Successfully loaded .env file")
}

// returns the value of the environment variable if it exists, otherwise returns an error
func GetEnv(key string, errorMessage string) (string, error) {
	value, ok := os.LookupEnv(key)

	if !ok {
		return "", errors.New(errorMessage)
	}

	return value, nil
}

// returns the value of the environment variable if it exists, otherwise returns the default value
func GetEnvWithDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	return value
}
