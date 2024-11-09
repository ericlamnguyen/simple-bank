package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// The struct field values are read from a config file or environment variables
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddr          string        `mapstructure:"SERVER_ADDR"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DRUATION"`
}

// LoadConfig reads configuration file or environment variables
func LoadConfig(path string) (config Config, err error) {
	// Specify location of the config file
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Read environment variables, existing env variables will take precedence over those defined in config file
	// This will also transform the key in app.env to match the format used for env variables, eg database_url will be transformed into DATABASE_URL
	viper.AutomaticEnv()

	// This function reads the configuration file into Viper's configuration management system.
	// After calling this, Viper will have access to the values defined in the file and can use them when you retrieve configuration settings.
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Store env variables in struct config
	err = viper.Unmarshal(&config)
	return
}
