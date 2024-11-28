package officer

import (
	"net/http"
	"web-http/utils"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/", officerHandler)
	r.Get("/dashboard", officerDashboardHandler)
	r.Get("/settings", officerSettingsHandler)
	r.Get("/books/{title}/page/{page}", officerBookPageHandler)

	return r
}

func officerHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	utils.SendResponse(res, "Go to /officer/dashboard or /officer/settings, please.", http.StatusNotFound, nil)
}

func officerDashboardHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	utils.SendResponse(res, "Welcome to Officer Dashboard!", http.StatusOK, nil)
}

func officerSettingsHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	utils.SendResponse(res, "Welcome to Officer Settings!", http.StatusOK, nil)
}

func officerBookPageHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	vars := map[string]string{
		"title": chi.URLParam(req, "title"),
		"page": chi.URLParam(req, "page"),
	}

	utils.SendResponse(res, "Welcome to Officer Book Page!", http.StatusOK, vars)
}