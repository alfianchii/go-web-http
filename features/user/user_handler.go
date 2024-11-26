package user

import (
	"net/http"
	"web-http/utils"

	"github.com/gorilla/mux"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	utils.ResponseWithName(res, req, "WELCOME to Go web development")
}

func AboutHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	utils.ResponseWithName(res, req, "You are on About's Go web development")
}

func GreetHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	vars := mux.Vars(req)
	name := vars["name"]
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