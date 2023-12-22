package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	envFilename := ".env"
	if v := os.Getenv("ENV_FILE"); v != "" {
		if _, err := os.Stat(v); err == nil {
			log.Fatalln(err)
		}
		envFilename = v
	}
	if envFilename == ".env" {
		if _, err := os.Stat(".env"); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return
			}
			log.Fatalln(err)
		}
	}

	_ = godotenv.Load(envFilename)
}

func Env(key, defaultValue string) string {
	if os.Getenv(key) == "" {
		return defaultValue
	}
	return os.Getenv(key)
}

func MustEnv(key string) string {
	if os.Getenv(key) == "" {
		log.Panicf("unknown %s param in ENV", key)
	}
	return os.Getenv(key)
}
