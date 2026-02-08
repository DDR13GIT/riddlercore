package config

import (
	"log"

	"github.com/spf13/viper"
)

// environment mode
const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

// Init initiates of config load
func Init() {
	viper.SetEnvPrefix("riddlercore")
	viper.AutomaticEnv() // Automatically bind all environment variables

	// Set local file details
	viper.SetConfigName("config") // Name of the file (without extension)
	viper.SetConfigType("yml")    // File extension
	viper.AddConfigPath(".")      // Look for config in the working directory

	// Read the local file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file 'config.yml' not found in root directory")
		} else {
			log.Fatalf("Error reading config file: %s", err)
		}
	}

	loadApp()
	loadDB()
	loadRedis()
	loadMQCfg()
}
