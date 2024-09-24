package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type HomePageData struct {
	Title           string
	IsAuthenticated bool
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Println("Ошибка парсинга шаблона:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Проверяем, есть ли токен (куки)
	_, err = r.Cookie("token")
	isAuthenticated := err == nil

	data := HomePageData{
		Title:           "Главная",
		IsAuthenticated: isAuthenticated,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Ошибка выполнения шаблона:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
