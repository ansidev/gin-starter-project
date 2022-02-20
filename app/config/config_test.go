package config

import (
	"github.com/ansidev/gin-starter-project/constant"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestConfig(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

type ConfigTestSuite struct {
	suite.Suite
}

func (s *ConfigTestSuite) SetupTest() {
	os.Unsetenv("APP_ENV")
	os.Unsetenv("LOG_LEVEL")
	os.Unsetenv("HOST")
	os.Unsetenv("PORT")
	os.Unsetenv("DB_DRIVER")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_USERNAME")
	os.Unsetenv("DB_PASSWORD")
}

func (s *ConfigTestSuite) TestLoadConfigProd() {
	os.Setenv("APP_ENV", constant.DefaultProdEnv)
	os.Setenv("LOG_LEVEL", "error")
	os.Setenv("HOST", "https://github.com")
	os.Setenv("PORT", "80")
	os.Setenv("DB_DRIVER", "postgres")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "demo")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")

	var config Config
	LoadConfig("..", "app.env.example", &config)

	assert.Equal(s.T(), os.Getenv("APP_ENV"), constant.DefaultProdEnv, "APP_ENV should be prod")
	assert.Equal(s.T(), "error", config.LogLevel, "LogLevel should be error")

	assert.Equal(s.T(), "https://github.com", config.Host, "Host should be https://github.com")
	assert.Equal(s.T(), 80, config.Port, "Port should be 80")

	assert.Equal(s.T(), "postgres", config.DbDriver, "DbHost should be postgres")
	assert.Equal(s.T(), "127.0.0.1", config.DbHost, "DbHost should be 127.0.0.1")
	assert.Equal(s.T(), 5432, config.DbPort, "DbPort should be 5432")
	assert.Equal(s.T(), "demo", config.DbName, "DbName should be demo")
	assert.Equal(s.T(), "postgres", config.DbUsername, "DbUsername should be postgres")
	assert.Equal(s.T(), "postgres", config.DbPassword, "DbPassword should be postgres")
}

func (s *ConfigTestSuite) TestLoadConfigFromEnvFile() {
	os.Setenv("APP_ENV", "local")

	var config Config
	LoadConfig("..", "app.env.example", &config)

	assert.Equal(s.T(), os.Getenv("APP_ENV"), "local", "APP_ENV should be local")
	assert.Equal(s.T(), "debug", config.LogLevel, "LogLevel should be debug")

	assert.Equal(s.T(), "localhost", config.Host, "Host should be localhost")
	assert.Equal(s.T(), 8080, config.Port, "Port should be 8080")

	assert.Equal(s.T(), "postgres", config.DbDriver, "DbHost should be postgres")
	assert.Equal(s.T(), "127.0.0.1", config.DbHost, "DbHost should be 127.0.0.1")
	assert.Equal(s.T(), 5432, config.DbPort, "DbPort should be 5432")
	assert.Equal(s.T(), "demo", config.DbName, "DbName should be demo")
	assert.Equal(s.T(), "postgres", config.DbUsername, "DbUsername should be postgres")
	assert.Equal(s.T(), "postgres", config.DbPassword, "DbPassword should be postgres")
}
