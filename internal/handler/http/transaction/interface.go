package auth

import "net/http"

type Handler interface {
	DownloadTemplate(w http.ResponseWriter, r *http.Request)
	Upload(w http.ResponseWriter, r *http.Request)
}
