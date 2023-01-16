package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var envVars *Environments

type Environments struct {
	DBName   string `mapstructure:"DATABASE_NAME"`
	DBSchema string `mapstructure:"DATABASE_SCHEMA"`
	UserName string `mapstructure:"DATABASE_USERNAME"`
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     string `mapstructure:"DATABASE_PORT"`
	Password string `mapstructure:"DATABASE_PASSWORD"`
	Charset  string `mapstructure:"DATABASE_CHARSET"`
}

func LoadEnvVars() *Environments {
	viper.SetConfigFile(".env")
	viper.SetDefault("API_PORT", "8080")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Sprintf("unable find or read configuration file: %w", err)
	}

	if err := viper.Unmarshal(&envVars); err != nil {
		fmt.Sprintf("unable to unmarshal configurations from environment: %w", err)
	}
	return envVars
}
