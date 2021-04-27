package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Schema struct
type Schema struct {
	Env          string `mapstructure:"env"`
	Port         int    `mapstructure:"port"`
	DefaultLimit int    `mapstructure:"default_limit"`
	MaxLimit     int    `mapstructure:"max_limit"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Options  string `mapstructure:"options"`
	} `mapstructure:"database"`
}

// Config global parameter config
var Config Schema

func init() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(".")             // Look for config in current directory
	config.AddConfigPath("config/")       // Optionally look for config in the working directory.
	config.AddConfigPath("../config/")    // Look for config needed for tests.
	config.AddConfigPath("../")           // Look for config needed for tests.
	config.AddConfigPath("../../config/") // Look for config needed for tests.
	config.AddConfigPath("../../")        // Look for config needed for tests.

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	config.AutomaticEnv()

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = config.Unmarshal(&Config)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	// fmt.Printf("Current Config: %+v", Config)
}

func GetConfig() *Schema {
	return &Config
}
