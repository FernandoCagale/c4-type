package routers

import (
	"github.com/FernandoCagale/c4-type/api/handlers"
	"github.com/gorilla/mux"
)

type SystemRoutes struct {
	healthHandler   *handlers.HealthHandler
	typeHandler *handlers.TypeHandler
}

func (routes *SystemRoutes) MakeHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", routes.healthHandler.Health).Methods("GET")
	r.HandleFunc("/types", routes.typeHandler.Create).Methods("POST")
	r.HandleFunc("/types", routes.typeHandler.FindAll).Methods("GET")
	r.HandleFunc("/types/{id}", routes.typeHandler.FindById).Methods("GET")
	r.HandleFunc("/types/{id}", routes.typeHandler.DeleteById).Methods("DELETE")

	return r
}

func NewSystem(healthHandler *handlers.HealthHandler, typeHandler *handlers.TypeHandler) *SystemRoutes {
	return &SystemRoutes{
		healthHandler:   healthHandler,
		typeHandler: typeHandler,
	}
}
