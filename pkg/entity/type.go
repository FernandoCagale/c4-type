package entity

import (
	"github.com/go-ozzo/ozzo-validation"
)

type Type struct {
	Code  string   `json:"code"`
	Types []string `json:"types"`
}

func (e Type) Validate() error {
	return validation.ValidateStruct(&e, )
}
