package websocket

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func BroadcastChats() {
	for {
		clientChat := <-BroadcastChan
		BroadcastTypers()

		if clientChat.Send {
			ClientsLock.Lock()
			chat := SetChat(clientChat)

			for client := range Clients {
				clientChat.CreatedAt = time.Now()
				err := client.WriteJSON(clientChat)
				if err != nil {
					log.Printf("Error writing chat: %v\n", err)
					client.Close()
					delete(Clients, client)
				}
			}

			InsertChat(chat)
			ClientsLock.Unlock()
		}
	}
}

func BroadcastTypers() {
	ClientsLock.Lock()
	defer ClientsLock.Unlock()

	typerList := []string{}
	for typer := range Typers {
		typerList = append(typerList, typer)
	}

	for client := range Clients {
		err := client.WriteJSON(map[string]interface{}{
			"typers": typerList,
		})
		if err != nil {
			log.Printf("Error writing typers: %v\n", err)
			client.Close()
			delete(Clients, client)
		}
	}
}

func CleanUpDisconnectedClients(conn *websocket.Conn) {
	ClientsLock.Lock()
	tabID := Clients[conn]
	delete(Clients, conn)
	delete(Typers, tabID)
	ClientsLock.Unlock()
	conn.Close()

	BroadcastTypers()
}