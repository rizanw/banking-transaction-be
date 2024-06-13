package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server OK!")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	return router
}
