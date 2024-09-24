package routes

import (
	"CarSharing/controllers"
	"CarSharing/middleware"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/").Subrouter()
	userRouter.Use(middleware.AuthMiddleware)

	userRouter.HandleFunc("/logout", controllers.Logout).Methods("GET")
	userRouter.HandleFunc("/profile", controllers.ProfilePage).Methods("GET")
	userRouter.HandleFunc("/cars", controllers.GetAllCars).Methods("GET")
	userRouter.HandleFunc("/cars/{id}", controllers.GetCar).Methods("GET")
	userRouter.HandleFunc("/bookings", controllers.GetUserBookings).Methods("GET")
	userRouter.HandleFunc("/bookings", controllers.CreateBooking).Methods("POST")
}
