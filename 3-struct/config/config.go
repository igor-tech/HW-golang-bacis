package config

import "os"

type Config struct {
	Key string
}

func NewConfig(env string) *Config {
	key := os.Getenv(env)
	return &Config{
		Key: key,
	}
}

func (key *Config) GetKey() string {
	return key.Key
}
