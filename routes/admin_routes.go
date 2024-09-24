// routes/admin_routes.go

package routes

import (
	"CarSharing/controllers"
	"CarSharing/middleware"

	"github.com/gorilla/mux"
)

func RegisterAdminRoutes(router *mux.Router) {
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.AdminMiddleware)

	// Админская панель
	adminRouter.HandleFunc("/dashboard", controllers.AdminDashboard).Methods("GET")

	// Управление пользователями
	adminRouter.HandleFunc("/users", controllers.ManageUsers).Methods("GET")
	adminRouter.HandleFunc("/users/{id}/make_admin", controllers.MakeAdmin).Methods("POST")

	// Управление автомобилями
	adminRouter.HandleFunc("/cars", controllers.ManageCars).Methods("GET")
	adminRouter.HandleFunc("/cars/add", controllers.AddCar).Methods("GET", "POST")
	adminRouter.HandleFunc("/cars/{id}/edit", controllers.EditCar).Methods("GET", "POST")
	adminRouter.HandleFunc("/cars/{id}/delete", controllers.DeleteCar).Methods("POST")
	adminRouter.HandleFunc("/booked_cars", controllers.ShowBookedCars).Methods("GET")
}
