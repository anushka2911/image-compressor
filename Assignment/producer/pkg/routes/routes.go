package routes

import (
	controllers "github.com/anushka/producer/pkg/controller"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {

	r.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
}
