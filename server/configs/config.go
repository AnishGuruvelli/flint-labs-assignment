package configs

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port   string
	Secret string
	DbURL  string
}

var Cfg *Config

func bindEnvVariables() {
	viper.BindEnv("env", "ENV")
}

func NewConfig() (*Config, error) {
	ctxlog := logrus.WithFields(logrus.Fields{"service": "NewConfig"})

	// Load ENV from .env file
	if err := godotenv.Load(); err != nil {
		errorMessage := "error loading .env file: " + err.Error()
		ctxlog.Error(errorMessage)
		return nil, errors.New(errorMessage)
	}

	defaultConfigName := "config"

	// Setting up default configs
	viper.SetConfigName(defaultConfigName)
	viper.AddConfigPath("configs")
	viper.SetConfigType("yaml")

	// Read default configs
	if err := viper.ReadInConfig(); err != nil {
		errorMessage := "error reading viper config: " + err.Error()
		ctxlog.Error(errorMessage)
		return nil, errors.New(errorMessage)
	}

	ctxlog.Info("Reading profile: " + os.Getenv("ENV"))

	// Resolving env and setting env-specific config
	if env := strings.ToLower(os.Getenv("ENV")); strings.Compare(env, "") != 0 {
		envConfigName := defaultConfigName + "-" + env
		viper.SetConfigName(envConfigName)
	}

	// Merging env configs
	if err := viper.MergeInConfig(); err != nil {
		errorMessage := "error merging env-specific config: " + err.Error()
		ctxlog.Error(errorMessage)
		return nil, errors.New(errorMessage)
	}

	bindEnvVariables()

	Cfg = &Config{}

	if err := viper.Unmarshal(Cfg); err != nil {
		errorMessage := "error converting config to struct: " + err.Error()
		ctxlog.Error(errorMessage)
		return nil, errors.New(errorMessage)
	}

	ctxlog.Info("Viper AllKeys:", viper.AllKeys())

	return Cfg, nil
}
