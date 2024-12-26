package user

import (
	"fmt"
	"net/http"
	"time"
	"web-http/config"

	"go.mongodb.org/mongo-driver/bson"
)

var collectionName string = "users"

type User struct {
	UserId string `bson:"userId,omitempty" json:"userId"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Email string `bson:"email" json:"email"`
	IsOnline bool `bson:"isOnline" json:"isOnline"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type UserResponse struct {
	UserId string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	IsOnline bool `json:"isOnline"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func CreateUser(user User) (UserResponse, error) {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	var userResponse UserResponse = UserResponse{
		UserId: user.UserId,
		Username: user.Username,
		Email: user.Email,
		IsOnline: user.IsOnline,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	
	_, err := config.MongoDB.Collection(collectionName).InsertOne(ctx, user)
	if err != nil {
		return userResponse, err
	}

	return userResponse, nil
}

func GetUserByUsername(username string) (User, error) {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	var user User
	err := config.MongoDB.Collection(collectionName).FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, fmt.Errorf("unauthorized; user not found")
	}

	return user, nil
}

func IsUsernameExist(username string) bool {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	isExist := config.MongoDB.Collection(collectionName).FindOne(ctx, bson.M{"username": username}).Err()

	return isExist == nil
}

func IsEmailExist(email string) bool {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	isExist := config.MongoDB.Collection(collectionName).FindOne(ctx, bson.M{"email": email}).Err()

	return isExist == nil
}

func SetUserOnline(username string) error {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	_, err := config.MongoDB.Collection(collectionName).UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"isOnline": true}})
	if err != nil {
		http.Error(nil, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func SetUserOffline(username string) error {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	_, err := config.MongoDB.Collection(collectionName).UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"isOnline": false}})
	if err != nil {
		http.Error(nil, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}