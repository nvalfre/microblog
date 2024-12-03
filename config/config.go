package config

import (
	"fmt"
	"log"

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
func LoadConfig(path string, env string) (*AppConfig, error) {
	var config AppConfig

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading default config file: %v", err)
		return nil, err
	}

	if env != "" {
		viper.SetConfigName(fmt.Sprintf("config.%s", env))
		if err := viper.MergeInConfig(); err != nil {
			log.Printf("Error reading %s config file: %v", env, err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Error unmarshalling config: %v", err)
		return nil, err
	}

	return &config, nil
}
