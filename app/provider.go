package app

import (
	"go.mongodb.org/mongo-driver/mongo"
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
	logger := logger.New(env, logger.FileTemplate("gocrawler-web-%Y_%m_%d"))
	return logger
}

// ProvideMongodb is function to init mongodb connection
func ProvideMongodb(conf *configuration.AppConfig) *mongo.Database {
	return db.NewMongodb(conf)
}
