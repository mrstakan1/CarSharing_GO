package controllers

import (
	"CarSharing/database"
	"CarSharing/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"html/template"
	"log"
	"net/http"
)

type ProfilePageData struct {
	Title    string
	User     models.User
	Bookings []models.Booking
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		log.Println("Ошибка парсинга шаблона:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

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

	var user models.User
	database.DB.First(&user, claims.UserID)

	var bookings []models.Booking
	database.DB.Preload("Car").Where("user_id = ?", claims.UserID).Find(&bookings)

	data := ProfilePageData{
		Title:    "Профиль пользователя",
		User:     user,
		Bookings: bookings,
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
		return JwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("некорректный токен")
	}

	return claims, nil
}
