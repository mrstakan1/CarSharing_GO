// controllers/admin_controller.go

package controllers

import (
	"github.com/gorilla/mux"

	"html/template"
	"log"
	"net/http"
	"strconv"

	"CarSharing/database"
	"CarSharing/models"
)

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
	}{
		Title: "Админ-панель",
	}

	tmpl.Execute(w, data)
}

func ManageUsers(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/manage_users.html")
	if err != nil {
		log.Println("Ошибка парсинга шаблона manage_users.html:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println("Ошибка получения пользователей:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Users []models.User
	}{
		Title: "Управление Пользователями",
		Users: users,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Ошибка выполнения шаблона manage_users.html:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func MakeAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["id"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Println("Некорректный ID пользователя:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		log.Println("Пользователь не найден:", result.Error)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	user.IsAdmin = true
	result = database.DB.Save(&user)
	if result.Error != nil {
		log.Println("Ошибка обновления пользователя:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/users", http.StatusSeeOther)
}

func ManageCars(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/admin/manage_cars.html")
	if err != nil {
		log.Println("Ошибка парсинга шаблона manage_cars.html:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var cars []models.Car
	result := database.DB.Find(&cars)
	if result.Error != nil {
		log.Println("Ошибка получения автомобилей:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Cars  []models.Car
	}{
		Title: "Управление Автомобилями",
		Cars:  cars,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Ошибка выполнения шаблона manage_cars.html:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/admin/add_car.html")
		if err != nil {
			log.Println("Ошибка парсинга шаблона add_car.html:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := struct {
			Title string
		}{
			Title: "Добавить Автомобиль",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println("Ошибка выполнения шаблона add_car.html:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Обработка POST-запроса для добавления автомобиля
	name := r.FormValue("make")
	model := r.FormValue("model")
	yearStr := r.FormValue("year")
	location := r.FormValue("location")
	availableStr := r.FormValue("available")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		log.Println("Некорректный год выпуска:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	available := false
	if availableStr == "on" {
		available = true
	}

	car := models.Car{
		Make:      name,
		Model:     model,
		Year:      year,
		Location:  location,
		Available: available,
	}

	result := database.DB.Create(&car)
	if result.Error != nil {
		log.Println("Ошибка добавления автомобиля:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/cars", http.StatusSeeOther)
}

func EditCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carIDStr := vars["id"]
	carID, err := strconv.Atoi(carIDStr)
	if err != nil {
		log.Println("Некорректный ID автомобиля:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var car models.Car
	result := database.DB.First(&car, carID)
	if result.Error != nil {
		log.Println("Автомобиль не найден:", result.Error)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/admin/edit_car.html")
		if err != nil {
			log.Println("Ошибка парсинга шаблона edit_car.html:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := struct {
			Title string
			Car   models.Car
		}{
			Title: "Редактировать Автомобиль",
			Car:   car,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println("Ошибка выполнения шаблона edit_car.html:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Обработка POST-запроса для редактирования автомобиля
	name := r.FormValue("make")
	model := r.FormValue("model")
	yearStr := r.FormValue("year")
	location := r.FormValue("location")
	availableStr := r.FormValue("available")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		log.Println("Некорректный год выпуска:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	available := false
	if availableStr == "on" {
		available = true
	}

	car.Make = name
	car.Model = model
	car.Year = year
	car.Location = location
	car.Available = available

	result = database.DB.Save(&car)
	if result.Error != nil {
		log.Println("Ошибка обновления автомобиля:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/cars", http.StatusSeeOther)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carIDStr := vars["id"]
	carID, err := strconv.Atoi(carIDStr)
	if err != nil {
		log.Println("Некорректный ID автомобиля:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var car models.Car
	result := database.DB.First(&car, carID)
	if result.Error != nil {
		log.Println("Автомобиль не найден:", result.Error)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	result = database.DB.Delete(&car)
	if result.Error != nil {
		log.Println("Ошибка удаления автомобиля:", result.Error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/cars", http.StatusSeeOther)
}
