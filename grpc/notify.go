package grpc

import (
	"context"
	"fmt"
	"github.com/FernandoCagale/c4-type/pkg/domain/types"
)

type ServerGRPC struct {
	usecase types.UseCase
}

func NewNotifyGRPC(usecase types.UseCase) *ServerGRPC {
	return &ServerGRPC{
		usecase: usecase,
	}
}

func (server *ServerGRPC) GetNotify(ctx context.Context, in *Request) (*Reply, error) {
	notify, err := server.usecase.FindById("001")
	if err != nil {
		fmt.Println(err)
	}
	return &Reply{Types: notify.Types}, nil
}
