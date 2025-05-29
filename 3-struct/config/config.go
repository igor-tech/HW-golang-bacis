package config

import "os"

type Config struct {
	key string
}

func NewConfig(env string) *Config {
	key := os.Getenv(env)
	return &Config{
		key: key,
	}
}
