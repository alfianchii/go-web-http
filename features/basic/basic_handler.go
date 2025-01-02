package basic

import (
	"html/template"
	"net/http"
	"web-http/dto"
	"web-http/utils"

	"github.com/go-chi/chi/v5"
)

type ContactDetails struct {
	Email string `json:"email"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	funcMap := template.FuncMap{
		"toUpper": utils.ToUpper,
	}
	
	view := template.Must(template.New("base.html").Funcs(funcMap).ParseFiles(
		"templates/layouts/base.html",
		"templates/pages/home.html",
	))

	data := dto.PageData{
		Title: "Home",
		Heading: "Welcome to Go web development",
		Content: "This is a simple web application using Go programming language.",
	}

	err := view.Execute(res, data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AboutHandler(res http.ResponseWriter, req *http.Request) {
	session, _ := utils.Store.Get(req, utils.SessionName)
	username := session.Values["username"].(string)

	utils.SendResponse(res, "Welcome to the About page, " + username, http.StatusOK, nil)
}

func AboutEmailHandler(res http.ResponseWriter, req *http.Request) {

	details := ContactDetails{
		Email: req.FormValue("email"),
		Subject: req.FormValue("subject"),
		Content: req.FormValue("content"),
	}

	utils.SendResponse(res, "Success get all Master Satker", http.StatusOK, details)
}

func GreetHandler(res http.ResponseWriter, req *http.Request) {

	name := chi.URLParam(req, "name")
	utils.SendResponse(res, "Hello, " + name + "!", http.StatusOK, nil)
}

func SearchHandler(res http.ResponseWriter, req *http.Request) {
	
	query := req.URL.Query().Get("q")
	if query == "" {
		utils.SendResponse(res, "No search query found.", http.StatusOK, nil)
		return
	}

	utils.SendResponse(res, "Search query is: " + query, http.StatusOK, nil)
}