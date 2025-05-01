package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env         string
	Port        string
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	env, err := getEnv("ENV")
	if err != nil {
		return nil, err
	}

	port, err := getEnv("PORT")
	if err != nil {
		return nil, err
	}

	databaseUrl, err := getEnv("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	return &Config{
		Env:         env,
		Port:        port,
		DatabaseURL: databaseUrl,
	}, nil
}

func getEnv(key string) (string, error) {
	if val := os.Getenv(key); val != "" {
		return val, nil
	}
	return "", fmt.Errorf("environment variable %s is not set", key)
}
