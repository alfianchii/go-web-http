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
	client, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v\n", err)
		return
	}
	defer CleanUpDisconnectedClients(client)

	ClientsLock.Lock()
	Clients[client] = ""
	ClientsLock.Unlock()

	messages := GetMessages()
	client.WriteJSON(messages)

	for {
		var clientMessage ClientMessage
		err := client.ReadJSON(&clientMessage)
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			break
		}

		ClientsLock.Lock()
		Clients[client] = clientMessage.ClientID
		ClientsLock.Unlock()

		if clientMessage.Typing && clientMessage.Text != "" {
			Typers[clientMessage.ClientID] = true
		} else if clientMessage.ClientID != "" {
			delete(Typers, clientMessage.ClientID)
		}

		BroadcastChan <-clientMessage
	}
}