package admin

import (
	"net/http"
	"web-http/utils"

	"github.com/go-chi/chi/v5"
)

func AdminHandler(res http.ResponseWriter, req *http.Request) {
	utils.SendResponse(res, "Go to /admin/dashboard or /admin/settings, please.", http.StatusNotFound, nil)
}

func AdminDashboardHandler(res http.ResponseWriter, req *http.Request) {
	utils.SendResponse(res, "Welcome to Admin Dashboard!", http.StatusOK, nil)
}

func AdminSettingsHandler(res http.ResponseWriter, req *http.Request) {
	utils.SendResponse(res, "Welcome to Admin Settings!", http.StatusOK, nil)
}

func AdminBookPageHandler(res http.ResponseWriter, req *http.Request) {
	vars := map[string]string{
		"title": chi.URLParam(req, "title"),
		"page": chi.URLParam(req, "page"),
	}
	
	utils.SendResponse(res, "Welcome to Admin Book Page!", http.StatusOK, vars)	
}