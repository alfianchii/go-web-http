package user

import (
	"context"
	"net/http"
	"time"
	"web-http/config"
	"web-http/utils"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
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

	collection := config.MongoDB.Collection("users")
	exist := collection.FindOne(context.TODO(), bson.M{"username": user.Username})
	if exist.Err() == nil {
		utils.SendResponse(res, "Username already exists", http.StatusBadRequest, nil)
		return
	}
	
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		utils.SendResponse(res, "Error creating user", http.StatusInternalServerError, nil)
		return
	}

	utils.SendResponse(res, "User created successfully", http.StatusOK, UserResponse{
		UserId: user.UserId,
		Username: user.Username,
		Email: user.Email,
		IsOnline: user.IsOnline,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}