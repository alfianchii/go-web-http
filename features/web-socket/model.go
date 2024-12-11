package websocket

import "github.com/gorilla/websocket"

type Message struct {
	ClientID  string `json:"clientId"`
	Text   string `json:"text,omitempty"`
	Send   bool   `json:"send,omitempty"`
	Typing bool   `json:"typing,omitempty"`
}

type Client struct {
	ID   string
	Conn *websocket.Conn
}