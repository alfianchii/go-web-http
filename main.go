package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"web-http/config"
	"web-http/features/admin"
	"web-http/features/auth"
	"web-http/features/basic"
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
	router.Use(middleware.CORS)
	
	utils.FileServer(router, "/public", http.Dir("./assets"))
	utils.FileServer(router, "/templates", http.Dir("./templates"))
	
	router.Get("/chats", websocket.WebsocketHandler)
	go websocket.BroadcastMessages()

	router.With(middleware.JWTMiddleware).Get("/", basic.HomeHandler)
	router.With(middleware.GuestMiddleware).Get("/register", auth.RegisterViewHandler)
	router.With(middleware.GuestMiddleware).Route("/login", func(r chi.Router) {
		r.Get("/", auth.LoginViewHandler)
		r.Post("/", auth.LoginHandler)
	})
	router.Post("/validate-jwt", auth.ValidateJWTHandler)
	router.Post("/logout", auth.LogoutHandler)
	router.With(middleware.AuthMiddleware).Get("/about", basic.AboutHandler)
	router.Post("/about", basic.AboutEmailHandler)
	router.Get("/greet/{name}", basic.GreetHandler)
	router.Get("/search", basic.SearchHandler)

	router.Route("/admin", func(r chi.Router) {
		r.Get("/", admin.AdminHandler)
		r.Get("/dashboard", admin.AdminDashboardHandler)
		r.Get("/settings", admin.AdminSettingsHandler)
		r.Get("/books/{title}/page/{page}", admin.AdminBookPageHandler)
		r.Get("/satker", satker.Handler(db).SatkerHandler)
	})

	router.Route("/user", func(r chi.Router) {
		r.Post("/", user.CreateUserHandler)
	})

	router.Mount("/officer", officer.Router())

	fmt.Printf("Server is running on http://%s\n", config.Address)
	log.Fatal(http.ListenAndServe(config.Address, router))
}