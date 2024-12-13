package user

import (
	"html/template"
	"net/http"
	"web-http/config"
	"web-http/utils"

	"github.com/go-chi/chi/v5"
)

type PageData struct {
	Title string
	Heading string
	Content string
}

type ContactDetails struct {
	Email string `json:"email"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Id string `json:"id"`
	Cookie string `json:"cookie"`
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	funcMap := template.FuncMap{
		"toUpper": utils.ToUpper,
	}
	
	view := template.Must(template.New("base.html").Funcs(funcMap).ParseFiles(
		"templates/layouts/base.html",
		"templates/pages/home.html",
	))

	data := PageData{
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

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)
	session, _ := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))

	creds := UserLogin{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	session.Values["username"] = creds.Username
	id := creds.Username + "123" + creds.Password
	session.Values["id"] = id
	session.Save(req, res)
	cookieValue := utils.CookieToString(res)

	utils.SendResponse(res, "You are logged in as "+creds.Username, http.StatusOK, UserLoginResponse{
		Username: creds.Username,
		Id: id,
		Cookie: cookieValue,
	})
}


func LogoutHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)
	session, _ := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))
	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1
	utils.Store.MaxAge(-1)
	session.Save(req, res)

	utils.SendResponse(res, "You are logged out", http.StatusOK, nil)
}

func AboutHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)
	session, _ := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))
	username := session.Values["username"].(string)

	utils.SendResponse(res, "Welcome to the About page, " + username, http.StatusOK, nil)
}

func AboutEmailHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	details := ContactDetails{
		Email: req.FormValue("email"),
		Subject: req.FormValue("subject"),
		Content: req.FormValue("content"),
	}

	utils.SendResponse(res, "Success get all Master Satker", http.StatusOK, details)
}

func GreetHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	name := chi.URLParam(req, "name")
	utils.SendResponse(res, "Hello, " + name + "!", http.StatusOK, nil)
}

func SearchHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)
	
	query := req.URL.Query().Get("q")
	if query == "" {
		utils.SendResponse(res, "No search query found.", http.StatusOK, nil)
		return
	}

	utils.SendResponse(res, "Search query is: " + query, http.StatusOK, nil)
}