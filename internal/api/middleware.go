package api

import (
	"github.com/alexgaas/order-reward/internal/usecase/auth"
	"net/http"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Get access token from request
		authParam := r.Header.Get("Authorization")

		// Check token to valid
		login, err := auth.ValidateToken(authParam)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}

		r.Header["Login"] = []string{login}

		// Token is authenticated, pass it through
		next.ServeHTTP(rw, r)
	})
}
