package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

// LoadConfig defines a method to load viper configs to memory
func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc/default")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Fatal("config file not found: ", err)
		}
	}
}
