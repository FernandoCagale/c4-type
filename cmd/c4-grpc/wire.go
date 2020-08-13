//+build wireinject

package main

import (
	"github.com/FernandoCagale/c4-type/grpc"
	"github.com/google/wire"
)

func SetupApplication() (*grpc.ServerGRPC, error) {
	wire.Build(grpc.Set)
	return nil, nil
}
