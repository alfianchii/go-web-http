package utils

import (
	"encoding/json"
	"net/http"
	"web-http/dto"

	"github.com/go-chi/chi/v5"
)

func SetHeaderJson(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
}

func ResponseSetup(res http.ResponseWriter) {
	SetHeaderJson(res)
}

func SendResponse(res http.ResponseWriter, msg string, status int, data interface{}) {
	ResponseSetup(res)

	response := dto.Response{
		Message: msg,
		Status: status,
		Data: data,
	}

	res.WriteHeader(status)
	json.NewEncoder(res).Encode(response)	
}

func FileServer(router chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		router.Get(path, http.RedirectHandler(path + "/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	router.Handle(path + "*", http.StripPrefix(path, http.FileServer(root)))
}