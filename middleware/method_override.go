package middleware

import (
	"net/http"
)

func MethodOverride(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err == nil {
				if method := r.FormValue("_method"); method != "" {
					r.Method = method
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
