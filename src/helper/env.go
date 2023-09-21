package helper

import "os"

func LoadEnv(envName string, def string) string {
	env := os.Getenv(envName)
	if env == "" {
		return def
	}

	return env
}
