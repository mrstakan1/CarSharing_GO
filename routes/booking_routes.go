package routes

import (
	"CarSharing/controllers"
	"CarSharing/middleware"

	"github.com/gorilla/mux"
)

func RegisterBookingRoutes(router *mux.Router) {
	bookingRouter := router.PathPrefix("/bookings").Subrouter()
	bookingRouter.Use(middleware.AuthMiddleware)

	bookingRouter.HandleFunc("", controllers.GetUserBookings).Methods("GET")
	bookingRouter.HandleFunc("", controllers.CreateBooking).Methods("POST")
}
