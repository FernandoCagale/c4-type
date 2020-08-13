package types

import (
	"github.com/FernandoCagale/c4-type/internal/errors"
	"github.com/FernandoCagale/c4-type/pkg/entity"
)

type InMemoryRepository struct {
	m map[string]*entity.Type
}

func NewInMemoryRepository() *InMemoryRepository {
	database := make(map[string]*entity.Type)
	database["001"] = &entity.Type{
		Code:  "001",
		Types: []string{"SMS", "EMAIL"},
	}
	return &InMemoryRepository{database}
}

func (repo *InMemoryRepository) FindAll() (result []*entity.Type, err error) {
	for _, types := range repo.m {
		result = append(result, types)
	}
	return result, nil
}

func (repo *InMemoryRepository) FindById(ID string) (types *entity.Type, err error) {
	for _, types := range repo.m {
		if types.Code == ID {
			return types, nil
		}
	}
	return nil, errors.ErrNotFound
}

func (repo *InMemoryRepository) DeleteById(ID string) (err error) {
	for _, types := range repo.m {
		if types.Code == ID {
			delete(repo.m, ID)
			return nil
		}
	}
	return errors.ErrNotFound
}

func (repo *InMemoryRepository) Create(e *entity.Type) (err error) {
	types := repo.m[e.Code]

	if types == nil {
		repo.m[e.Code] = e
		return nil
	}

	return nil
}
