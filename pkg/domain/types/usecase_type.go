package types

import (
	"github.com/FernandoCagale/c4-type/internal/errors"
	"github.com/FernandoCagale/c4-type/pkg/entity"
)

type TypeUseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *TypeUseCase {
	return &TypeUseCase{
		repo: repo,
	}
}

func (usecase *TypeUseCase) FindAll() (types []*entity.Type, err error) {
	return usecase.repo.FindAll()
}

func (usecase *TypeUseCase) FindById(ID string) (types *entity.Type, err error) {
	return usecase.repo.FindById(ID)
}

func (usecase *TypeUseCase) DeleteById(ID string) (err error) {
	return usecase.repo.DeleteById(ID)
}

func (usecase *TypeUseCase) Create(e *entity.Type) error {
	err := e.Validate()
	if err != nil {
		return errors.ErrInvalidPayload
	}

	if err = usecase.repo.Create(e); err != nil {
		return err
	}

	return nil
}
