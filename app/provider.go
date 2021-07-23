package app

import (
	"database/sql"
	"gocrawler-web-sample/infrastructure/configuration"
	"gocrawler-web-sample/infrastructure/db"
	"gocrawler-web-sample/infrastructure/environment"
	"gocrawler-web-sample/infrastructure/logger"

	"github.com/sirupsen/logrus"
)

// ProvideAppEnvironment is a function to provide app enviroment data
func ProvideAppEnvironment() (environment.AppEnvironment, error) {
	return environment.FromOsEnv()
}

// ProvideAppEnvConfig is a function to get AppConfig struct data
func ProvideAppEnvConfig(conf *configuration.AppConfig) configuration.AppConfig {
	return *conf
}

// ProvideLogger is a function to log http request and deployment to a file
func ProvideLogger(env environment.AppEnvironment) *logrus.Logger {
	logger := logger.New(env, logger.FileTemplate("command-product-app-%Y_%m_%d"))
	return logger
}

// ProvidePostgres is function to init postgres connection
func ProvidePostgres(config *configuration.AppConfig) *sql.DB {
	return db.NewPostgres(config)
}
