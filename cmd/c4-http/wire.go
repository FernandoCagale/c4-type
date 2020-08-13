//+build wireinject

package main

import (
	"github.com/FernandoCagale/c4-type/api/routers"
	"github.com/FernandoCagale/c4-type/pkg"
	"github.com/google/wire"
)

func SetupApplication() (*routers.SystemRoutes, error) {
	wire.Build(pkg.Container)
	return nil, nil
}