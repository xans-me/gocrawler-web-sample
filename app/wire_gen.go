// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package app

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"gocrawler-web-sample/infrastructure/configuration"
	"gocrawler-web-sample/infrastructure/environment"
)

// Injectors from configWire.go:

func InitializeAppConfig() (*configuration.AppConfig, error) {
	environmentConfigBinderProperties := _wireEnvironmentConfigBinderPropertiesValue
	appConfig := ProvideAppConfig(environmentConfigBinderProperties)
	return appConfig, nil
}

var (
	_wireEnvironmentConfigBinderPropertiesValue = configuration.EnvironmentConfigBinderProperties{
		FileName: "app-config",
		Path:     "./env",
	}
)

// Injectors from injector.go:

func InjectAppEnvironment() (environment.AppEnvironment, error) {
	appEnvironment, err := ProvideAppEnvironment()
	if err != nil {
		return "", err
	}
	return appEnvironment, nil
}

func InjectAppConfig() configuration.AppConfig {
	environmentConfigBinderProperties := _wireEnvironmentConfigBinderPropertiesValue
	appConfig := ProvideAppConfig(environmentConfigBinderProperties)
	configurationAppConfig := ProvideAppEnvConfig(appConfig)
	return configurationAppConfig
}

func InjectMongoDB() *mongo.Database {
	environmentConfigBinderProperties := _wireEnvironmentConfigBinderPropertiesValue
	appConfig := ProvideAppConfig(environmentConfigBinderProperties)
	database := ProvideMongodb(appConfig)
	return database
}

// configWire.go:

var (
	configModuleSets = wire.NewSet(wire.Value(configuration.EnvironmentConfigBinderProperties{
		FileName: "app-config",
		Path:     "./env",
	}), ProvideAppConfig)
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

// injector.go:

var (
	AppModule = wire.NewSet(
		configModuleSets,
		ProvideAppEnvironment,
		ProvideAppEnvConfig,
		ProvideMongodb,
	)
)
