package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketHandler(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v\n", err)
		return
	}
	defer CleanUpDisconnectedClients(conn)

	ClientsLock.Lock()
	Clients[conn] = ""
	ClientsLock.Unlock()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			break
		}

		ClientsLock.Lock()
		Clients[conn] = msg.ClientID
		ClientsLock.Unlock()

		if msg.Typing && msg.Text != "" {
			Typers[msg.ClientID] = true
		} else if msg.ClientID != "" {
			delete(Typers, msg.ClientID)
		}

		BroadcastChan <-msg
	}
}