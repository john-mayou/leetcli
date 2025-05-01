package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env                string
	Port               string
	DbURL              string
	RedisURL           string
	FrontendURL        string
	JWTSecret          string
	GithubClientID     string
	GithubClientSecret string
	GithubRedirectURI  string
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

	dbUrl, err := getEnv("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	redisUrl, err := getEnv("REDIS_URL")
	if err != nil {
		return nil, err
	}

	frontendUrl, err := getEnv("FRONTEND_URL")
	if err != nil {
		return nil, err
	}

	jwtSecret, err := getEnv("JWT_SECRET")
	if err != nil {
		return nil, err
	}

	githubClientID, err := getEnv("GITHUB_CLIENT_ID")
	if err != nil {
		return nil, err
	}

	githubClientSecret, err := getEnv("GITHUB_CLIENT_SECRET")
	if err != nil {
		return nil, err
	}

	githubRedirectURI, err := getEnv("GITHUB_REDIRECT_URI")
	if err != nil {
		return nil, err
	}

	return &Config{
		Env:                env,
		Port:               port,
		DbURL:              dbUrl,
		RedisURL:           redisUrl,
		FrontendURL:        frontendUrl,
		JWTSecret:          jwtSecret,
		GithubClientID:     githubClientID,
		GithubClientSecret: githubClientSecret,
		GithubRedirectURI:  githubRedirectURI,
	}, nil
}

func getEnv(key string) (string, error) {
	if val := os.Getenv(key); val != "" {
		return val, nil
	}
	return "", fmt.Errorf("environment variable %s is not set", key)
}
