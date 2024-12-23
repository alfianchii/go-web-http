package token

import (
	"log"
	"time"
	"web-http/config"

	"go.mongodb.org/mongo-driver/bson"
)

var collectionName string = "tokens"

type Token struct {
	Username string `bson:"username" json:"username"`
	Value string `bson:"value,omitempty" json:"value"`
	IsBlacklist bool `bson:"isBlacklist" json:"isBlacklist"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

func GetValidTokenFromUser(username string) (Token, error) {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	var token Token
	filters := bson.M{
		"username": username, 
		"isBlacklist": false,
	}
	err := config.MongoDB.Collection(collectionName).FindOne(ctx, filters).Decode(&token)
	if err != nil {
		log.Printf("Error getting token: %v\n", err)
		return token, err
	}

	return token, nil
}

func CreateToken(username string, jwtToken string) (Token, error) {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	var token Token = Token{
		Username: username,
		Value: jwtToken,
		IsBlacklist: false,
		CreatedAt: time.Now(),
	}
	
	_, err := config.MongoDB.Collection(collectionName).InsertOne(ctx, token)
	if err != nil {
		log.Printf("Error inserting token: %v\n", err)
		return token, err
	}

	return token, nil
}

func BlacklistUsedTokens(username string) error {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	filters := bson.M{
		"username": username,
		"isBlacklist": false,
	}
	_, err := config.MongoDB.Collection("tokens").UpdateMany(ctx, filters, bson.M{"$set": bson.M{"isBlacklist": true}})
	if err != nil {
		log.Printf("Error blacklisting tokens: %v\n", err)
		return err
	}

	return nil
}