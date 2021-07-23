package kurs

import (
	"gocrawler-web-sample/infrastructure/environment"
)

type Service struct {
	appEnvironment   environment.AppEnvironment
	publicRepository IKursRepository
}

// NewService function to init service instance
func NewService(appEnvironment environment.AppEnvironment, publicRepository IKursRepository) *Service {
	return &Service{appEnvironment: appEnvironment, publicRepository: publicRepository}
}
