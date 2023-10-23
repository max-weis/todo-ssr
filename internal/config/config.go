package config

import (
	"fmt"
	"os"
)

type Config struct {
	Production bool
	Port       string
}

func NewConfig() *Config {
	return &Config{
		Production: getEnvVarWithDefault("PROD", "false") == "true",
		Port:       getEnvVarWithDefault("PORT", "8080"),
	}
}

func getEnvVarWithDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (c *Config) HttpPort() string {
	return fmt.Sprintf(":%s", c.Port)
}
