package user

import (
	"net/http"
	"time"
	"web-http/config"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	UserId string `bson:"user_id,omitempty" json:"userId"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Email string `bson:"email" json:"email"`
	IsOnline bool `bson:"is_online" json:"isOnline"`
	CreatedAt time.Time `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time `bson:"updated_at" json:"updatedAt"`
}

type UserResponse struct {
	UserId string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	IsOnline bool `json:"isOnline"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func GetUserByUsername(username string) (User, error) {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	var user User
	err := config.MongoDB.Collection("users").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func SetUserOnline(username string) error {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	_, err := config.MongoDB.Collection("users").UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"is_online": true}})
	if err != nil {
		http.Error(nil, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}

func SetUserOffline(username string) error {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	_, err := config.MongoDB.Collection("users").UpdateOne(ctx, bson.M{"username": username}, bson.M{"$set": bson.M{"is_online": false}})
	if err != nil {
		http.Error(nil, err.Error(), http.StatusInternalServerError)
		return err
	}

	return nil
}