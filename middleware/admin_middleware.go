package middleware

import (
	"net/http"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAdmin := r.Context().Value("isAdmin")
		if isAdmin == nil || isAdmin.(bool) == false {
			http.Error(w, "Доступ запрещен", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
