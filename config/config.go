package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	DB
	Auth
}

type DB struct {
	Database string `env:"MYSQL_DATABASE"`
	UserName string `env:"MYSQL_USER"`
	Password string `env:"MYSQL_PASSWORD"`
	Host     string `env:"MYSQL_HOST"`
	Port     int    `env:"MYSQL_PORT"`
}

type Auth struct {
	AccessTokenExpirationHour int    `env:"ACCESS_TOKEN_EXPIRATION_HOUR"`
	AccessTokenSecret         string `env:"ACCESS_TOKEN_SECRET"`
}

func NewConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("failed to load .env file. %w", err)
	}

	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("filed to parse environment variables. %w", err)
	}

	return cfg, nil
}
