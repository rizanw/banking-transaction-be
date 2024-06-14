package auth

import "net/http"

type Handler interface {
	Register(w http.ResponseWriter, r *http.Request)
}
