package main

import (
	"CarSharing/database"
	"CarSharing/middleware"
	"CarSharing/routes"

	"github.com/gorilla/mux"

	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	database.Connect()

	router := mux.NewRouter()

	staticFiles := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticFiles))

	router.Use(middleware.AuthMiddleware)

	routes.RegisterPublicRoutes(router) // home, login, register, logout
	routes.RegisterUserRoutes(router)   // profile, booking
	routes.RegisterAdminRoutes(router)  // add/edit/delete car, admin-panel

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Сервер запущен на порту", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}

}
