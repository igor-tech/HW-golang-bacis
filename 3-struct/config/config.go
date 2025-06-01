package config

import (
	"fmt"
	"os"
)

type Config struct {
	Key string
}

func NewConfig(keyName string) (*Config, error) {
	key := os.Getenv(keyName)
	if key == "" {
		return nil, fmt.Errorf("environment variable %s is not set", keyName)
	}
	return &Config{
		Key: key,
	}, nil
}

func (key *Config) GetKey() string {
	return key.Key
}
