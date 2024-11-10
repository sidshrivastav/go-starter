package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config holds the loaded configuration values
type Config struct {
	AppName    string `mapstructure:"app.name"`
	AppVersion string `mapstructure:"app.version"`
	AppPort    int    `mapstructure:"app.port"`
	Database   struct {
		Host     string `mapstructure:"database.host"`
		Port     int    `mapstructure:"database.port"`
		User     string `mapstructure:"database.user"`
		Password string `mapstructure:"database.password"`
	} `mapstructure:"database"`
	Logging struct {
		Level string `mapstructure:"logging.level"`
	} `mapstructure:"logging"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("yaml")
	viper.AddConfigPath("config")

	// Determine the environment (e.g., "dev", "prod", "stage")
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // Default to 'dev' if no environment is set
	}

	// Load the base configuration file
	viper.SetConfigFile("./config/base.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading base config file: %w", err)
	}

	// Load the environment-specific configuration file (e.g., config.prod.yaml)
	envConfigFile := fmt.Sprintf("%s.yaml", env)
	viper.SetConfigFile(envConfigFile)
	if err := viper.MergeInConfig(); err != nil {
		return nil, fmt.Errorf("error reading environment config file %s: %w", envConfigFile, err)
	}

	// Automatically load environment variables (like DATABASE_HOST, LOGGING_LEVEL)
	viper.AutomaticEnv()

	// Optionally, you can read from a .env file in local development
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Warning: Could not read .env file. Make sure it exists if you're using one.")
	}

	// Unmarshal the configuration into a Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %w", err)
	}

	return &config, nil
}
