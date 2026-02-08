package config

import (
	"strings"

	"github.com/spf13/viper"
)

// scrts is the default application configuration
var scrts = map[string]string{}

// GetAuthSecret returns the application secret for id
func GetAuthSecret(id string) string {
	return scrts[strings.ToLower(id)]
}

// loadAuthSecrets loads application secrets
func loadAuthSecrets() {
	scrts = viper.GetStringMapString("auth_secrets")
}
