package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"web-http/config"
	"web-http/features/admin"
	"web-http/features/officer"
	"web-http/features/satker"
	"web-http/features/user"
	websocket "web-http/features/web-socket"
	"web-http/middleware"
	"web-http/utils"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	config.InitENV()
	
	db := config.DBConnect()
	defer db.Close()
	mongoClient := config.InitMongoDB()
	defer mongoClient.Disconnect(context.Background())
	utils.InitCookie()

	router := chi.NewRouter()

	router.Use(chiMiddleware.Logger)
	
	utils.FileServer(router, "/public", http.Dir("./assets"))
	utils.FileServer(router, "/templates", http.Dir("./templates"))
	
	router.Get("/chats", websocket.WebsocketHandler)
	go websocket.BroadcastMessages()

	router.Get("/", user.HomeHandler)
	router.Post("/login", user.LoginHandler)
	router.Post("/logout", user.LogoutHandler)
	router.With(middleware.AuthMiddleware).Get("/about", user.AboutHandler)
	router.Post("/about", user.AboutEmailHandler)
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