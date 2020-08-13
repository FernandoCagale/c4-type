package grpc

import (
	"github.com/FernandoCagale/c4-type/pkg/domain/types"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewNotifyGRPC, types.Set)
