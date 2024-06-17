package utils

import "net/http"

type Handler interface {
	GetCorporates(w http.ResponseWriter, r *http.Request)
}
