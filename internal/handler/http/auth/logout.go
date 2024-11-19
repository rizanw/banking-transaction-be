package auth

import "net/http"

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		err error
	)

	secretToken, ok := ctx.Value("token").(string)
	if secretToken == "" || !ok {
		http.Error(w, "user not login", http.StatusUnauthorized)
		return
	}

	err = h.ucAuth.Logout(secretToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
