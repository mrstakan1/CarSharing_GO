package middleware

import (
	"context"
	"net/http"
)

//var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем наличие токена (куки)
		_, err := r.Cookie("token")
		isAuthenticated := err == nil

		// Передаем информацию о том, авторизован ли пользователь, в контекст запроса
		ctx := r.Context()
		ctx = context.WithValue(ctx, "isAuthenticated", isAuthenticated)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//func AuthMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		cookie, err := r.Cookie("token")
//		if err != nil {
//			if err == http.ErrNoCookie {
//				http.Redirect(w, r, "/login", http.StatusSeeOther)
//				return
//			}
//
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//
//		tokenStr := cookie.Value
//		claims := &controllers.Claims{}
//
//		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
//			return jwtKey, nil
//		})
//
//		if err != nil || !token.Valid {
//			http.Redirect(w, r, "/login", http.StatusSeeOther)
//		}
//
//		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
//		ctx = context.WithValue(ctx, "isAdmin", claims.Admin)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
