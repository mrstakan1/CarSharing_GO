// controllers/dashboard_controller.go

package controllers

import (
	"CarSharing/database"
	"CarSharing/models"
	"html/template"
	"log"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		log.Println("Ошибка парсинга шаблона dashboard.html:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	userID, ok := r.Context().Value("userID").(uint)
	if !ok {
		log.Println("Не удалось получить userID из контекста")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		log.Println("Пользователь не найден:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Name  string
	}{
		Title: "Дашборд",
		Name:  user.Name,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Ошибка выполнения шаблона dashboard.html:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
