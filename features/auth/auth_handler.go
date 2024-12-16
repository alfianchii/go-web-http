package auth

import (
	"net/http"
	"web-http/config"
	"web-http/utils"
)


type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Id string `json:"id"`
	Cookie string `json:"cookie"`
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