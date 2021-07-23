package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gocrawler-web-sample/app"
	"gocrawler-web-sample/domain/kurs"
	"gocrawler-web-sample/infrastructure/configuration"

	log "github.com/sirupsen/logrus"
)

func main() {

	// initialize config
	config, err := app.InitializeAppConfig()
	if err != nil {
		panic(err.Error())
	}

	// set bugSnag
	log.Info("############################")
	log.Info("gocrawler-web-sample")
	log.Info("Powered by : Golang Web Crawler Example ", app.Version)
	log.Info("You app running on ", config.App.Environment, " mode")
	log.Info("############################")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// kurs routes
	kursRoutes, err := kurs.InjectRoutes()
	if err != nil {
		panic(err.Error())
	}
	kursRoutes.RegisterRoutes(router)

	configuration.Listen(*config, router)
}
