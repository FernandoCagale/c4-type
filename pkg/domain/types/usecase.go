package types

import (
	"github.com/FernandoCagale/c4-type/pkg/entity"
)

type UseCase interface {
	Create(types *entity.Type) (err error)
	FindAll() (types []*entity.Type, err error)
	FindById(ID string) (types *entity.Type, err error)
	DeleteById(ID string) (err error)
}
