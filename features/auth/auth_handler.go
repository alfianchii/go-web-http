package auth

import (
	"html/template"
	"net/http"
	"web-http/config"
	"web-http/dto"
	userModel "web-http/features/user"
	"web-http/utils"
)


type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

func LoginViewHandler(res http.ResponseWriter, req *http.Request) {
	view := template.Must(template.New("base.html").ParseFiles(
		"templates/layouts/base.html",
		"templates/pages/auth/login.html",
	))

	data := dto.PageData{
		Title: "Login",
		Heading: "Login",
		Content: "Please login to continue",
	}

	err := view.Execute(res, data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	session, _ := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))

	creds := UserLogin{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	if creds.Username == "" || creds.Password == "" {
		utils.SendResponse(res, "Please provide a username and password", http.StatusBadRequest, nil)
		return
	}

	user, err := userModel.GetUserByUsername(creds.Username)
	isPassword := utils.ComparePassword(user.Password, creds.Password)
	if isPassword != nil || err != nil {
		utils.SendResponse(res, "Username or password is invalid", http.StatusUnauthorized, nil)
		return
	}
	
	token, err := utils.GenerateJWT(creds.Username)
	if err != nil {
		utils.SendResponse(res, "Error generating token", http.StatusInternalServerError, nil)
		return
	}

	session.Values["username"] = creds.Username
	session.Values["token"] = token
	session.Save(req, res)
	userModel.SetUserOnline(creds.Username)

	utils.SendResponse(res, "You are logged in as "+creds.Username, http.StatusOK, UserLoginResponse{
		Username: creds.Username,
		Token: token,
	})
}

func LogoutHandler(res http.ResponseWriter, req *http.Request) {
	session, _ := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))
	username := session.Values["username"].(string)
	authHeader := req.Header.Get("Authorization")[len("Bearer "):]

	if len(authHeader) == 4 || session.Values["token"] == nil {
		utils.SendResponse(res, "Unauthorized; missing authorization token.", http.StatusUnauthorized, nil)
		return
	}

	claims, err := utils.ValidateJWT(authHeader)
	if err != nil {
		utils.SendResponse(res, "Unauthorized; invalid authorization token.", http.StatusUnauthorized, nil)
		return
	}

	utils.RemoveCookie(res, req, session)
	userModel.SetUserOffline(username)
	
	utils.SendResponse(res, "You are logged out", http.StatusOK, claims)
}

func ValidateJWTHandler(res http.ResponseWriter, req *http.Request) {
	session, _ := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))
	username := session.Values["username"].(string)
	authHeader := req.Header.Get("Authorization")[len("Bearer "):]
	
	if len(authHeader) == 4 || session.Values["token"] == nil {
		userModel.SetUserOffline(username)
		utils.RemoveCookie(res, req, session)
		utils.SendResponse(res, "Unauthorized; missing authorization token.", http.StatusUnauthorized, nil)
		return
	}

	claims, err := utils.ValidateJWT(authHeader)
	if err != nil {
		userModel.SetUserOffline(username)
		utils.RemoveCookie(res, req, session)
		utils.SendResponse(res, "Unauthorized; invalid authorization token.", http.StatusUnauthorized, nil)
		return
	}

	utils.SendResponse(res, "Valid token", http.StatusOK, claims)
}