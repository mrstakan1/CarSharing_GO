package middleware

import (
	"context"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("token")
		isAuthenticated := err == nil

		ctx := r.Context()
		ctx = context.WithValue(ctx, "isAuthenticated", isAuthenticated)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
