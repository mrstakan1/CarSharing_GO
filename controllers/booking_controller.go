package controllers

import (
	"CarSharing/database"
	"CarSharing/models"
	"html/template"
	"net/http"
	"time"
)

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	carID := r.FormValue("car_id")
	userID := r.Context().Value("UserID").(uint)

	var car models.Car
	database.DB.First(&car, carID)

	if car.Available == false {
		http.Error(w, "Автомобиль недоступен для бронирования", http.StatusBadRequest)
		return
	}

	booking := models.Booking{
		UserID:    userID,
		CarID:     car.ID,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(2 * time.Hour), // todo time length
	}
	database.DB.Create(&booking)

	car.Available = false
	database.DB.Save(&car)

	http.Redirect(w, r, "/bookings", http.StatusSeeOther)
}

func GetUserBookings(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uint)
	var bookings []models.Booking
	database.DB.Preload("Car").Where("user_id = ?", userID).Find(&bookings)

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/bookings.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title    string
		Bookings []models.Booking
	}{
		Title:    "Мои бронирования",
		Bookings: bookings,
	}

	tmpl.Execute(w, data) //todo err
}
