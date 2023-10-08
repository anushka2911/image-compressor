package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/anushka/producer/pkg/routes"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
