package config

import "os"

func GetEnvValueSafely(key string, errorMessage string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		panic(errorMessage)
	}

	return value
}
