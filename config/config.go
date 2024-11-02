package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type BaseConfig struct {
	Port     int    `yaml:"port"`
	LogLevel string `yaml:"log_level"`
}

type Config struct {
	BaseConfig
	// Add any environment-specific fields here if needed
	DatabaseURL string `yaml:"database_url"`
}

// LoadConfig reads the configuration based on the provided environment
func LoadConfig(env string) Config {
	var baseConfigFile = "config/base.yaml"
	var envConfigFile string
	switch env {
	case "dev":
		envConfigFile = "config/dev.yaml"
	case "stage":
		envConfigFile = "config/stage.yaml"
	case "prod":
		envConfigFile = "config/prod.yaml"
	default:
		log.Fatalf("Unknown environment %s", env)
	}

	baseConfig := loadBaseConfig(baseConfigFile)
	envConfig := loadEnvConfig(envConfigFile)

	return Config{
		BaseConfig:  baseConfig,
		DatabaseURL: envConfig.DatabaseURL,
	}
}

// Load base configuration from base config file
func loadBaseConfig(file string) BaseConfig {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read base config file: %v", err)
	}

	var baseConfig BaseConfig
	if err := yaml.Unmarshal(data, &baseConfig); err != nil {
		log.Fatalf("failed to unmarshal base config: %v", err)
	}

	return baseConfig
}

// Load environment-specific configuration
func loadEnvConfig(file string) Config {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to read env config file: %v", err)
	}

	var envConfig Config
	if err := yaml.Unmarshal(data, &envConfig); err != nil {
		log.Fatalf("failed to unmarshal env config: %v", err)
	}

	return envConfig
}
