package config

import (
	"os"

	"github.com/joho/godotenv"
)

type GoDotEnvProvider struct {
	configFolder string
	logger       logger
}

type logger interface {
	Log(args ...interface{})
	Logf(format string, a ...interface{})
	Warn(args ...interface{})
	Warnf(format string, a ...interface{})
	Error(args ...interface{})
	Errorf(format string, a ...interface{})
}

func NewGoDotEnvProvider(l logger, configFolder string) *GoDotEnvProvider {
	provider := &GoDotEnvProvider{
		configFolder: configFolder,
		logger:       l,
	}

	provider.readConfig(configFolder)

	return provider
}

func (g *GoDotEnvProvider) readConfig(confLocation string) {
	defaultFile := confLocation + "/.env"

	env := os.Getenv("GOFR_ENV")
	if env == "" {
		env = "local"
	}

	overrideFile := confLocation + "/." + env + ".env"

	err := godotenv.Load(overrideFile)
	if err == nil {
		g.logger.Log("Loaded config from file: ", overrideFile)
	}

	err = godotenv.Load(defaultFile)
	if err == nil {
		g.logger.Log("Loaded config from file: ", defaultFile)
	}
}

func (g *GoDotEnvProvider) Get(key string) string {
	return os.Getenv(key)
}

func (g *GoDotEnvProvider) GetOrDefault(key, defaultValue string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}

	return defaultValue
}
