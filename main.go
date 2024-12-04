package main

import (
	"fmt"
	"log"
	"net/http"
	"web-http/config"
	"web-http/features/admin"
	"web-http/features/officer"
	"web-http/features/satker"
	"web-http/features/user"
	"web-http/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := config.DBConnect()
	defer db.Close()

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	
	utils.FileServer(router, "/public", http.Dir("./assets"))
	
	router.Get("/", user.HomeHandler)
	router.Get("/about", user.AboutHandler)
	router.Get("/greet/{name}", user.GreetHandler)
	router.Get("/search", user.SearchHandler)

	router.Route("/admin", func(r chi.Router) {
		r.Get("/", admin.AdminHandler)
		r.Get("/dashboard", admin.AdminDashboardHandler)
		r.Get("/settings", admin.AdminSettingsHandler)
		r.Get("/books/{title}/page/{page}", admin.AdminBookPageHandler)
		r.Get("/satker", satker.Handler(db).SatkerHandler)
	})

	router.Mount("/officer", officer.Router())

	fmt.Printf("Server is running on http://%s\n", config.Address)
	log.Fatal(http.ListenAndServe(config.Address, router))
}