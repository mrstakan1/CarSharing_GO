// routes/public_routes.go

package routes

import (
	"CarSharing/controllers"

	"github.com/gorilla/mux"
)

func RegisterPublicRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.HomePage).Methods("GET")

	router.HandleFunc("/login", controllers.ShowLoginPage).Methods("GET")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	router.HandleFunc("/register", controllers.ShowRegisterPage).Methods("GET")
	router.HandleFunc("/register", controllers.Register).Methods("POST")

	router.HandleFunc("/logout", controllers.Logout).Methods("GET")
}
