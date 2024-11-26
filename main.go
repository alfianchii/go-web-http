package main

import (
	"fmt"
	"log"
	"net/http"
	"web-http/config"
	"web-http/features/admin"
	"web-http/features/user"
	"web-http/middleware"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/", user.HomeHandler).Methods("GET")
	router.HandleFunc("/about", user.AboutHandler).Methods("GET")
	router.HandleFunc("/greet/{name}", user.GreetHandler).Methods("GET")
	router.HandleFunc("/search", user.SearchHandler).Methods("GET")
	router.HandleFunc("/admin", admin.AdminHandler).Methods("GET")

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/dashboard", admin.AdminDashboardHandler).Methods("GET")
	adminRouter.HandleFunc("/settings", admin.AdminSettingsHandler).Methods("GET")
	adminRouter.HandleFunc("/books/{title}/page/{page}", admin.AdminBookPageHandler).Methods("GET")

	router.Use(middleware.Logging)

	fmt.Printf("Server is running on http://%s\n", config.Address)
	log.Fatal(http.ListenAndServe(config.Address, router))
}