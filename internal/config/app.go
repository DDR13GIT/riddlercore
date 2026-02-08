package config

import (
	"time"

	"github.com/spf13/viper"
)

// Version represents app version
var Version = "unversioned"

// Application holds the application configuration
type Application struct {
	Base            string
	Env             string
	HTTPPort        int
	Sentry          string
	Version         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	PaginationLimit int
}

// app is the default application configuration
var app Application

// App returns the default application configuration
func App() *Application {
	return &app
}

// loadApp loads application configuration
func loadApp() {
	env := EnvDevelopment
	if e := viper.GetString("app.env"); e != "" {
		env = e
	}
	app = Application{
		Base:            viper.GetString("app.base"),
		HTTPPort:        viper.GetInt("app.http_port"),
		Env:             env,
		Sentry:          viper.GetString("sentry.dsn"),
		Version:         Version,
		ReadTimeout:     viper.GetDuration("app.read_timeout") * time.Second,
		WriteTimeout:    viper.GetDuration("app.write_timeout") * time.Second,
		IdleTimeout:     viper.GetDuration("app.idle_timeout") * time.Second,
		PaginationLimit: viper.GetInt("app.pagination_limit"),
	}
}
