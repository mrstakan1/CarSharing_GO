package controllers

import (
	"CarSharing/database"
	"CarSharing/models"

	"github.com/dgrijalva/jwt-go"

	"errors"
	"html/template"
	"log"
	"net/http"
)

type ProfilePageData struct {
	Title string
	User  models.User
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		log.Println("Ошибка парсинга шаблона:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Получаем идентификатор пользователя из контекста или куки
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	claims, err := ParseToken(cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Получаем данные пользователя из базы данных
	var user models.User
	database.DB.First(&user, claims.UserID)

	data := ProfilePageData{
		Title: "Профиль пользователя",
		User:  user,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Ошибка выполнения шаблона:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("некорректный токен")
	}

	return claims, nil
}
