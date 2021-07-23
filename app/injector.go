//+build wireinject

package app

import (
	"database/sql"
	"github.com/google/wire"
	"gocrawler-web-sample/infrastructure/configuration"
	"gocrawler-web-sample/infrastructure/environment"
)

var (
	AppModule = wire.NewSet(
		configModuleSets,
		ProvideAppEnvironment,
		ProvideAppEnvConfig,
		ProvidePostgres,
	)
)

func InjectAppEnvironment() (environment.AppEnvironment, error) {
	panic(wire.Build(AppModule))
}

func InjectAppConfig() configuration.AppConfig {
	panic(wire.Build(AppModule))
}

func InjectPostgres() *sql.DB {
	panic(wire.Build(AppModule))
}
