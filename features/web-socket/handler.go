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

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
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
		Clients[conn] = msg.TabID
		ClientsLock.Unlock()

		if msg.Typing && msg.Text != "" {
			Typers[msg.TabID] = true
		} else if msg.TabID != "" {
			delete(Typers, msg.TabID)
		}

		BroadcastChan <-msg
	}
}