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
	router.Handle("/api/transactions",
		middleware.VerifyAuth(&conf.JWT, []int32{user.RoleMaker, user.RoleApprover},
			http.HandlerFunc(handlerTransaction.GetTransactions))).Methods(http.MethodGet)
	router.Handle("/api/transaction/{transactionID}",
		middleware.VerifyAuth(&conf.JWT, []int32{user.RoleMaker, user.RoleApprover},
			http.HandlerFunc(handlerTransaction.GetTransaction))).Methods(http.MethodGet)
	router.Handle("/api/transaction/{transactionID}/audit",
		middleware.VerifyAuth(&conf.JWT, []int32{user.RoleApprover},
			http.HandlerFunc(handlerTransaction.AuditTransaction))).Methods(http.MethodPost)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server OK!")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	return router
}
