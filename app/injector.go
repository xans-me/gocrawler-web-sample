//+build wireinject

package app

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"gocrawler-web-sample/infrastructure/configuration"
	"gocrawler-web-sample/infrastructure/environment"
)

var (
	AppModule = wire.NewSet(
		configModuleSets,
		ProvideAppEnvironment,
		ProvideAppEnvConfig,
		ProvideMongodb,
	)
)

func InjectAppEnvironment() (environment.AppEnvironment, error) {
	panic(wire.Build(AppModule))
}

func InjectAppConfig() configuration.AppConfig {
	panic(wire.Build(AppModule))
}

func InjectMongoDB() *mongo.Database {
	panic(wire.Build(AppModule))
}
