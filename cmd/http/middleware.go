package main

import (
	"context"
	"net/http"
	"tx-bank/internal/infra/session"
)

func MiddlewareAuth(jwt session.Manager, roles []int32, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get auth token from header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}
		tokenString = tokenString[len("Bearer "):]

		// Validate the JWT token
		userID, role, err := jwt.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// pass session to context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user_id", userID)
		ctx = context.WithValue(ctx, "role", role)
		ctx = context.WithValue(ctx, "token", tokenString)
		r = r.WithContext(ctx)

		// verify authorization
		if !isRoleAllowed(role, roles) {
			http.Error(w, "Forbidden Access", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isRoleAllowed(role int32, allowedRoles []int32) bool {
	for _, v := range allowedRoles {
		if v == role {
			return true
		}
	}
	return false
}
