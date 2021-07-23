// +build wireinject

package kurs

import (
	"github.com/google/wire"
	"gocrawler-web-sample/app"
)

var (
	ModuleSets = wire.NewSet(
		NewRepository,
		NewService,
		wire.Bind(new(IKursRepository), new(*Repository)),
		wire.Bind(new(IKursService), new(*Service)),
		NewDelivery,
		NewRoutes)
)

func InjectRoutes() (*Routes, error) {
	panic(wire.Build(app.AppModule, ModuleSets))
}
