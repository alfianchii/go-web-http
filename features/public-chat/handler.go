package websocket

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ChatsHandler(res http.ResponseWriter, req *http.Request) {
	client, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v\n", err)
		return
	}
	defer CleanUpDisconnectedClients(client)

	userAgent := req.UserAgent()
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	
	ClientsLock.Lock()
	Clients[client] = ""
	ClientsLock.Unlock()

	chats := GetChats()
	client.WriteJSON(chats)
	BroadcastTypers()

	for {
		var clientChat ClientChat
		err := client.ReadJSON(&clientChat)
		if err != nil {
			log.Printf("Error reading chat: %v\n", err)
			break
		}

		ClientsLock.Lock()
		Clients[client] = clientChat.Username
		ClientsLock.Unlock()

		if clientChat.Typing && clientChat.Text != "" {
			Typers[clientChat.Username] = true
		} else if clientChat.Username != "" {
			delete(Typers, clientChat.Username)
		}

		clientChat.UserAgent = userAgent
		clientChat.IPAddress = ip
		
		BroadcastChan <-clientChat
	}
}