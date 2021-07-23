package kurs

import (
	"gocrawler-web-sample/infrastructure/configuration"
)

type Repository struct {
	appConfig configuration.AppConfig
}

// NewRepository to create new repository instance
func NewRepository(appConfig configuration.AppConfig) *Repository {
	return &Repository{appConfig: appConfig}
}
