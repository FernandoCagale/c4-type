package pkg

import (
	"github.com/FernandoCagale/c4-type/api/handlers"
	"github.com/FernandoCagale/c4-type/api/routers"
	"github.com/FernandoCagale/c4-type/pkg/domain/types"
	"github.com/google/wire"
)

var Container = wire.NewSet(types.Set, handlers.Set, routers.Set)
