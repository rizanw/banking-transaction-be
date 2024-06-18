package auth

import "net/http"

type Handler interface {
	DownloadTemplate(w http.ResponseWriter, r *http.Request)
	Upload(w http.ResponseWriter, r *http.Request)
	GetTransactions(w http.ResponseWriter, r *http.Request)
	GetTransaction(w http.ResponseWriter, r *http.Request)
	AuditTransaction(w http.ResponseWriter, r *http.Request)
	GetTransactionStats(w http.ResponseWriter, r *http.Request)
}
