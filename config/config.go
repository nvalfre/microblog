package config

import (
	"fmt"
	"microblog/infrastructure/logger"
	"os"

	"github.com/spf13/viper"
)

// AppConfig holds all configuration constants
type AppConfig struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Auth     AuthConfig     `mapstructure:"auth"`
}

// ServerConfig holds server-related constants
type ServerConfig struct {
	Port string `mapstructure:"port"`
}

// DatabaseConfig holds MongoDB-related constants
type DatabaseConfig struct {
	URI      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}

// RedisConfig holds Redis-related constants
type RedisConfig struct {
	Address string `mapstructure:"address"`
}

// AuthConfig holds JWT-related constants
type AuthConfig struct {
	JWTSecret string `mapstructure:"jwt_secret"`
}

// LoadConfig reads and load configuration by environment
func LoadConfig(path string) (*AppConfig, error) {
	env := os.Getenv("APP_ENV")
	var config AppConfig

	if env != "" {
		viper.SetConfigName(fmt.Sprintf("config-%s", env))
	} else {
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Error reading default config file", err)
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		logger.Error("Error unmarshalling config", err)
		return nil, err
	}

	return &config, nil
}
