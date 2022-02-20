package config

import "github.com/ansidev/gin-starter-project/pkg/db"

var (
	AppConfig Config
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	Host           string `mapstructure:"HOST"`
	Port           int    `mapstructure:"PORT"`
	db.SqlDbConfig `mapstructure:",squash"`
}
