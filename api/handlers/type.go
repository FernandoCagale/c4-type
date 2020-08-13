package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/FernandoCagale/c4-type/api/render"
	"github.com/FernandoCagale/c4-type/internal/errors"
	"github.com/FernandoCagale/c4-type/pkg/domain/types"
	"github.com/FernandoCagale/c4-type/pkg/entity"
	"github.com/gorilla/mux"
	"net/http"
)

type TypeHandler struct {
	usecase types.UseCase
}

func NewTypeHandler(usecase types.UseCase) *TypeHandler {
	return &TypeHandler{
		usecase: usecase,
	}
}

func (handler *TypeHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	types, err := handler.usecase.FindAll()
	if err != nil {
		fmt.Println(err.Error())
		render.ResponseError(w, err, http.StatusInternalServerError)
		return
	}

	render.Response(w, types, http.StatusOK)
}

func (handler *TypeHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID := vars["id"]

	types, err := handler.usecase.FindById(ID)
	if err != nil {
		fmt.Println(err.Error())
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, types, http.StatusOK)
}

func (handler *TypeHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID := vars["id"]

	err := handler.usecase.DeleteById(ID)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusNoContent)
}

func (handler *TypeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var types *entity.Type

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&types); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := handler.usecase.Create(types); err != nil {
		switch err {
		case errors.ErrInvalidPayload:
			render.ResponseError(w, err, http.StatusBadRequest)
		case errors.ErrConflict:
			render.ResponseError(w, err, http.StatusConflict)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusCreated)
}
