package auth

import (
	"html/template"
	"net/http"
	"time"
	"web-http/config"
	"web-http/dto"
	tokenModel "web-http/features/token"
	userModel "web-http/features/user"
	"web-http/utils"

	"github.com/gorilla/sessions"
)

func RegisterViewHandler(res http.ResponseWriter, req *http.Request) {
	view := template.Must(template.New("base.html").ParseFiles(
		"templates/layouts/base.html",
		"templates/pages/auth/register.html",
	))

	data := dto.PageData{
		Title: "Register",
		Heading: "Register",
		Content: "Let's make an account for you",
	}

	err := view.Execute(res, data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
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

	err = tokenModel.BlacklistUsedTokens(creds.Username)
	if err != nil {
		utils.SendResponse(res, "Error blacklisting tokens", http.StatusInternalServerError, err)
		return
	}
	
	storedToken, err := tokenModel.CreateToken(creds.Username, token)
	if err != nil {
		utils.SendResponse(res, "Error storing token", http.StatusInternalServerError, err)
		return
	}

	session, _ := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))
	session.Values["username"] = storedToken.Username
	session.Values["token"] = storedToken.Value
	session.Values["exp"] = time.Now().Add(config.TokenDuration).Unix()
	session.Save(req, res)

	userModel.SetUserOnline(storedToken.Username)

	utils.SendResponse(res, "You are logged in as "+storedToken.Username, http.StatusOK, UserLoginResponse{
		Username: storedToken.Username,
		Token: token,
	})
}

func LogoutHandler(res http.ResponseWriter, req *http.Request) {
	session, err := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))
	if err != nil {
		utils.SendResponse(res, "Unauthorized; session is missing.", http.StatusUnauthorized, nil)
		return
	}

	sessionExp, ok := session.Values["exp"].(int64)
	if !ok || time.Now().Unix() > sessionExp {
		utils.SendResponse(res, "Unauthorized; session expired.", http.StatusUnauthorized, nil)
		return
	}

	sessionUsername := session.Values["username"].(string)
	sessionToken := session.Values["token"]

	authHeader, err := utils.GetBearerToken(req.Header.Get("Authorization"))
	if err != nil {
		utils.SendResponse(res, utils.CapitalizeFirstLetter(err.Error()), http.StatusUnauthorized, nil)
		return
	}

	token, err := tokenModel.GetValidTokenFromUser(sessionUsername)
	validToken := token.Value
	if err != nil {
		utils.SendResponse(res, "Unauthorized; token was blacklisted.", http.StatusUnauthorized, nil)
		return
	}

	if validToken != authHeader || validToken != sessionToken {
		utils.SendResponse(res, "Unauthorized; invalid authorization token.", http.StatusUnauthorized, nil)
		return
	}

	claims, err := utils.ValidateJWT(validToken)
	if err != nil {
		utils.SendResponse(res, "Unauthorized; token was expired.", http.StatusUnauthorized, nil)
		return
	}

	invalidateUser(res, req, session)
	
	utils.SendResponse(res, "You are logged out", http.StatusOK, claims)
}

func ValidateJWTHandler(res http.ResponseWriter, req *http.Request) {
	session, err := utils.Store.Get(req, config.GetENV("COOKIE_NAME"))
	if err != nil {
		utils.SendResponse(res, "Unauthorized; session is missing.", http.StatusUnauthorized, nil)
		return
	}

	sessionExp, ok := session.Values["exp"].(int64)
	if !ok || time.Now().Unix() > sessionExp {
		invalidateUser(res, req, session)
		utils.SendResponse(res, "Unauthorized; session expired.", http.StatusUnauthorized, nil)
		return
	}

	sessionUsername := session.Values["username"].(string)
	sessionToken := session.Values["token"]
	
	authHeader, err := utils.GetBearerToken(req.Header.Get("Authorization"))
	if err != nil {
		utils.SendResponse(res, utils.CapitalizeFirstLetter(err.Error()), http.StatusUnauthorized, nil)
		return
	}

	token, err := tokenModel.GetValidTokenFromUser(sessionUsername)
	validToken := token.Value
	if err != nil {
		utils.SendResponse(res, "Unauthorized; token was blacklisted.", http.StatusUnauthorized, nil)
		return
	}

	if validToken != authHeader || validToken != sessionToken {
		invalidateUser(res, req, session)
		utils.SendResponse(res, "Unauthorized; invalid authorization token.", http.StatusUnauthorized, nil)
		return
	}

	claims, err := utils.ValidateJWT(validToken)
	if err != nil {
		invalidateUser(res, req, session)
		utils.SendResponse(res, "Unauthorized; token was expired.", http.StatusUnauthorized, nil)
		return
	}

	_, err = userModel.GetUserByUsername(claims.Username)
	if err != nil {
		invalidateUser(res, req, session)
		utils.SendResponse(res, "Unauthorized; user not found.", http.StatusUnauthorized, nil)
		return
	}

	var response ValidTokenResponse = ValidTokenResponse{
		Username: claims.Username,
		Token: validToken,
		Exp: claims.ExpiresAt.Time.Unix(),
	}
	utils.SendResponse(res, "Valid token", http.StatusOK, response)
}

func invalidateUser(res http.ResponseWriter, req *http.Request, session *sessions.Session) {
	username := session.Values["username"].(string)
	
	userModel.SetUserOffline(username)
	utils.RemoveCookie(res, req, session)
	tokenModel.BlacklistUsedTokens(username)
}