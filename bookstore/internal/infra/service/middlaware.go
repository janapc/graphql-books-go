package service

import (
	"net/http"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "invalid token", http.StatusForbidden)
			return
		}
		auth = auth[len("Bearer "):]
		_, err := ValidateToken(auth)
		if err != nil {
			http.Error(w, "invalid token", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
