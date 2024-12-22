package user

import (
	"net/http"
	"time"
	"web-http/utils"

	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(res http.ResponseWriter, req *http.Request) {
	user := User{
		UserId: ksuid.New().String(),
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
		Email: req.FormValue("email"),
		IsOnline: false,
		CreatedAt: time.Now(),
	}

	if user.Username == "" || user.Password == "" || user.Email == "" {
		utils.SendResponse(res, "Please provide a username, password, and email", http.StatusBadRequest, nil)
		return
	}

	existUsername := IsUsernameExist(user.Username)
	if existUsername {
		utils.SendResponse(res, "Username already exists", http.StatusBadRequest, nil)
		return
	}
	
	existEmail := IsEmailExist(user.Email)
	if existEmail {
		utils.SendResponse(res, "Email already exists", http.StatusBadRequest, nil)
		return
	}
	
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	createdUser, err := CreateUser(user)
	if err != nil {
		utils.SendResponse(res, "Error creating user", http.StatusInternalServerError, err)
		return
	}

	utils.SendResponse(res, "User created successfully", http.StatusOK, createdUser)
}