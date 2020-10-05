// Package conf contains a collection of structs and methods related to
//  initializing and loading the app's base configuration variables.
package conf

import (
	"github.com/nkarpenko/koho-transaction/common/model"
	"github.com/spf13/viper"
)

// Config of the service.
type Config struct {
	Name       string        `mapstructure:"name"`
	Desc       string        `mapstructure:"desc"`
	InputFile  string        `mapstructure:"input"`
	OutputFile string        `mapstructure:"output"`
	Limits     *model.Limits `mapstructure:"limits"`
	Version    string        `mapstructure:"version"`
}

// Load the config file
func Load(file string) (*Config, error) {
	var config *Config

	// Set config file type to yml and define default file.
	viper.SetConfigType("yml")
	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.SetConfigName("config")
	}

	// Read and load the config vars.
	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	// Parse config into struct.
	config = new(Config)
	if err := viper.Unmarshal(config); err != nil {
		return config, err
	}

	return config, nil
}
