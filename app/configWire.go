//+build wireinject

package app

import (
	"github.com/google/wire"
	"gocrawler-web-sample/infrastructure/configuration"
)

var (
	configModuleSets = wire.NewSet(wire.Value(configuration.EnvironmentConfigBinderProperties{
		FileName: "app-config",
		Path:     "./env",
	}),
		ProvideAppConfig)
)

func ProvideAppConfig(properties configuration.EnvironmentConfigBinderProperties) *configuration.AppConfig {
	environmentConfigBinder := configuration.NewEnvironmentConfigBinder(properties)
	environmentConfigBinder.Bind()
	config, err := environmentConfigBinder.GetAppConfig()
	if err != nil {
		panic(err)
	}
	return config
}

func InitializeAppConfig() (*configuration.AppConfig, error) {
	panic(wire.Build(configModuleSets))
}
