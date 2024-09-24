package controllers

import (
	"CarSharing/database"
	"CarSharing/models"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	var cars []models.Car
	database.DB.Find(&cars)

	isAuthenticated := r.Context().Value("userID") != nil

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/cars.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title           string
		Cars            []models.Car
		IsAuthenticated bool
	}{
		Title:           "Доступные автомобили",
		Cars:            cars,
		IsAuthenticated: isAuthenticated,
	}

	tmpl.Execute(w, data)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var car models.Car
	database.DB.First(&car, params["id"])

	isAuthenticated := r.Context().Value("userID") != nil

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/car_detail.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title           string
		Car             models.Car
		IsAuthenticated bool
	}{
		Title:           fmt.Sprintf("%s %s", car.Make, car.Model),
		Car:             car,
		IsAuthenticated: isAuthenticated,
	}

	tmpl.Execute(w, data)
}
