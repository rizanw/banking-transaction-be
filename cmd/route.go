package main

import (
	"log"
	"net/http"
	hAuth "tx-bank/internal/handler/http/auth/module"

	"github.com/gorilla/mux"
)

func newRoutes(uc UseCase) *mux.Router {
	router := mux.NewRouter()

	handlerAuth := hAuth.New(uc.Auth)
	router.HandleFunc("/api/register", handlerAuth.Register).Methods(http.MethodPost)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server OK!")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	return router
}
