package websocket

import (
	"context"
	"log"
	"time"
	"web-http/config"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName string = "chats"

type ClientChat struct {
	Username string `json:"username"`
	Text string `json:"text,omitempty"`
	Send bool `json:"send,omitempty"`
	Typing bool `json:"typing,omitempty"`
	UserAgent string `json:"userAgent,omitempty"`
	IPAddress string `json:"ipAddress,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Token string `json:"token"`
}

type Chat struct {
	Username string `bson:"username" json:"username"`
	Text  string `bson:"text" json:"text"`
	UserAgent string `bson:"userAgent" json:"userAgent"`
	IPAddress string `bson:"ipAddress" json:"ipAddress"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

type Client struct {
	ID   string
	Conn *websocket.Conn
}

func SetChat(clientChat ClientChat) Chat {
	return Chat{
		Username: clientChat.Username,
		Text: clientChat.Text,
		UserAgent: clientChat.UserAgent,
		IPAddress: clientChat.IPAddress,
		CreatedAt: time.Now(),
	}
}

func InsertChat(chat Chat) {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	_, err := config.MongoDB.Collection(collectionName).InsertOne(ctx, chat)
	if err != nil {
		log.Printf("Error inserting chat: %v\n", err)
	}
}

func GetChats() []Chat {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	findOptions := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	cursor, err := config.MongoDB.Collection(collectionName).Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Printf("Error getting chats: %v\n", err)
	}

	var chats []Chat
	if err = cursor.All(config.CtxBg(), &chats); err != nil {
		log.Printf("Error decoding chats: %v\n", err)
	}
	
	return chats
}