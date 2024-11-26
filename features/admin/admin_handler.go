package admin

import (
	"net/http"
	"web-http/utils"

	"github.com/gorilla/mux"
)

func AdminHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	utils.SendResponse(res, "Go to /admin/dashboard or /admin/settings, please.", http.StatusNotFound, nil)
}

func AdminDashboardHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	utils.SendResponse(res, "Welcome to Admin Dashboard!", http.StatusOK, nil)
}

func AdminSettingsHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	utils.SendResponse(res, "Welcome to Admin Settings!", http.StatusOK, nil)
}

func AdminBookPageHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseSetup(res, req)

	vars := mux.Vars(req)
	utils.SendResponse(res, "Welcome to Admin Book Page!", http.StatusOK, vars)	
}