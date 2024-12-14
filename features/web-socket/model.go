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

var collectionName string = "messages"

type ClientMessage struct {
	ClientID  string `json:"clientId"`
	Text   string `json:"text,omitempty"`
	Send   bool   `json:"send,omitempty"`
	Typing bool   `json:"typing,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

type Message struct {
	ClientID string `bson:"clientId" json:"clientId"`
	Text  string `bson:"text" json:"text"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

type Client struct {
	ID   string
	Conn *websocket.Conn
}

func SetMessage(clientMessage ClientMessage) Message {
	return Message{
		ClientID: clientMessage.ClientID,
		Text: clientMessage.Text,
		CreatedAt: time.Now(),
	}
}

func InsertMessage(message Message) {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	_, err := config.MongoDB.Collection(collectionName).InsertOne(ctx, message)
	if err != nil {
		log.Printf("Error inserting message: %v\n", err)
	}
}

func GetMessages() []Message {
	var ctx, cancel = config.CtxTime()
	defer cancel()

	findOptions := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}})
	cursor, err := config.MongoDB.Collection(collectionName).Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Printf("Error getting messages: %v\n", err)
	}

	var messages []Message
	if err = cursor.All(context.Background(), &messages); err != nil {
		log.Printf("Error decoding messages: %v\n", err)
	}
	
	return messages
}