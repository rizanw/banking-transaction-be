package main

import (
	"log"
	"net/http"
	"tx-bank/internal/common/middleware"
	"tx-bank/internal/config"
	hAuth "tx-bank/internal/handler/http/auth/module"
	hTransaction "tx-bank/internal/handler/http/transaction/module"
	"tx-bank/internal/model/user"

	"github.com/gorilla/mux"
)

func newRoutes(uc UseCase, conf *config.Config) *mux.Router {
	router := mux.NewRouter()

	handlerAuth := hAuth.New(uc.Auth)
	router.HandleFunc("/api/register", handlerAuth.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/login", handlerAuth.Login).Methods(http.MethodPost)

	handlerTransaction := hTransaction.New(uc.Transaction)
	router.Handle("/api/transaction/download-template",
		middleware.VerifyAuth(&conf.JWT, []int32{user.RoleMaker},
			http.HandlerFunc(handlerTransaction.DownloadTemplate))).Methods(http.MethodGet)
	router.Handle("/api/transaction/upload",
		middleware.VerifyAuth(&conf.JWT, []int32{user.RoleMaker},
			http.HandlerFunc(handlerTransaction.Upload))).Methods(http.MethodPost)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server OK!")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	return router
}
