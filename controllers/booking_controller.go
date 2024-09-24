package controllers

import (
	"CarSharing/database"
	"CarSharing/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type BookingPageData struct {
	Title   string
	Cars    []models.Car
	Message string // Сообщение для пользователя
}

func ShowBookingPage(w http.ResponseWriter, r *http.Request) {
	// Получаем список доступных автомобилей (только те, которые доступны и не забронированы)
	var cars []models.Car
	database.DB.Where("available = ?", true).Find(&cars)

	data := BookingPageData{
		Title: "Бронирование автомобиля",
		Cars:  cars,
	}

	tmpl, err := template.ParseFiles("templates/booking.html")
	if err != nil {
		log.Println("Ошибка парсинга шаблона:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Ошибка выполнения шаблона:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func BookCar(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из куки (JWT токен)
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

	carIDStr := r.FormValue("car_id")
	carID, err := strconv.ParseUint(carIDStr, 10, 32)
	if err != nil {
		log.Println("Ошибка преобразования carID:", err)
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	hoursStr := r.FormValue("hours")
	hours, err := strconv.Atoi(hoursStr)
	if err != nil || hours < 1 || hours > 8 {
		log.Println("Неверное количество часов:", err)
		http.Error(w, "Invalid hours value", http.StatusBadRequest)
		return
	}

	startTime := time.Now()

	endTime := startTime.Add(time.Duration(hours) * time.Hour)

	booking := models.Booking{
		UserID:    claims.UserID,
		CarID:     uint(carID),
		StartTime: startTime,
		EndTime:   endTime,
		CreatedAt: time.Now(),
	}

	var car models.Car
	database.DB.First(&car, carID)
	car.Available = false
	database.DB.Save(&car)

	database.DB.Create(&booking)

	data := BookingPageData{
		Title:   "Бронирование автомобиля",
		Cars:    []models.Car{}, // Пустой список, так как автомобиль уже забронирован
		Message: "Вы успешно забронировали автомобиль " + car.Make + " " + car.Model + " на " + strconv.Itoa(hours) + " часов.",
	}

	tmpl, err := template.ParseFiles("templates/booking.html")
	if err != nil {
		log.Println("Ошибка парсинга шаблона:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Ошибка выполнения шаблона:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
