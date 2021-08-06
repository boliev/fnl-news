package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
}

// NewConfig constructor
func NewConfig() (*Config, error) {
	if _, err := os.Stat("config.yaml"); err == nil {
		viper.SetConfigName("config.yaml")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {
			return nil, fmt.Errorf("fatal error config file: %s ", err.Error())
		}
	}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	err := viper.MergeInConfig() // Find and read the config file
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %s ", err.Error())
	}

	if _, err := os.Stat("./.env.local"); err == nil {
		viper.SetConfigName(".env.local")
		err = viper.MergeInConfig() // Find and read the config file
		if err != nil {
			return nil, fmt.Errorf("fatal error config file: %s ", err.Error())
		}
	}

	return &Config{}, nil
}

// GetString returns string from config
func (c Config) GetString(key string) string {
	return viper.GetString(key)
}

// GetInt returns int from config
func (c Config) GetInt(key string) int {
	return viper.GetInt(key)
}

// UnmarshalKey fill struct from config
func (c Config) UnmarshalKey(key string, strct interface{}) error {
	return viper.UnmarshalKey(key, strct)
}
