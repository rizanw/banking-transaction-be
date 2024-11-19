package main

import (
	"log"
	"net/http"
	"tx-bank/internal/domain/user"
	hauth "tx-bank/internal/handler/http/auth"
	hUtils "tx-bank/internal/handler/http/corporate"
	htransaction "tx-bank/internal/handler/http/transaction"
	"tx-bank/internal/infra/session"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func newRoutes(uc UseCase, session session.Manager) http.Handler {
	router := mux.NewRouter()

	handlerCorporate := hUtils.New(uc.Corporate)
	router.HandleFunc("/api/corporates", handlerCorporate.GetCorporates).Methods(http.MethodGet)

	handlerAuth := hauth.New(uc.Auth)
	router.HandleFunc("/api/send-otp", handlerAuth.SendOTP).Methods(http.MethodPost)
	router.HandleFunc("/api/register", handlerAuth.Register).Methods(http.MethodPost)
	router.HandleFunc("/api/login", handlerAuth.Login).Methods(http.MethodPost)
	router.Handle("/api/logout",
		MiddlewareAuth(session, []int32{int32(user.RoleMaker), int32(user.RoleApprover)},
			http.HandlerFunc(handlerAuth.Logout))).Methods(http.MethodPost)

	handlerTransaction := htransaction.New(uc.Transaction)
	router.Handle("/api/transaction/download-template",
		MiddlewareAuth(session, []int32{int32(user.RoleMaker)},
			http.HandlerFunc(handlerTransaction.DownloadTemplate))).Methods(http.MethodGet)
	router.Handle("/api/transaction/upload",
		MiddlewareAuth(session, []int32{int32(user.RoleMaker)},
			http.HandlerFunc(handlerTransaction.Upload))).Methods(http.MethodPost)
	router.Handle("/api/transactions",
		MiddlewareAuth(session, []int32{int32(user.RoleMaker), int32(user.RoleApprover)},
			http.HandlerFunc(handlerTransaction.GetTransactions))).Methods(http.MethodGet)
	router.Handle("/api/transactions/stats",
		MiddlewareAuth(session, []int32{int32(user.RoleMaker), int32(user.RoleApprover)},
			http.HandlerFunc(handlerTransaction.GetTransactionStats))).Methods(http.MethodGet)
	router.Handle("/api/transaction/{transactionID}",
		MiddlewareAuth(session, []int32{int32(user.RoleMaker), int32(user.RoleApprover)},
			http.HandlerFunc(handlerTransaction.GetTransaction))).Methods(http.MethodGet)
	router.Handle("/api/transaction/{transactionID}/audit",
		MiddlewareAuth(session, []int32{int32(user.RoleApprover)},
			http.HandlerFunc(handlerTransaction.AuditTransaction))).Methods(http.MethodPost)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("server OK!")
		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)

	// cors config
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	return handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)
}
