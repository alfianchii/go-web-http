package websocket

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func BroadcastMessages() {
	for {
		clientMessage := <-BroadcastChan
		BroadcastTypers()

		if clientMessage.Send {
			ClientsLock.Lock()
			message := SetMessage(clientMessage)

			for client := range Clients {
				clientMessage.CreatedAt = time.Now()
				err := client.WriteJSON(clientMessage)
				if err != nil {
					log.Printf("Error writing message: %v\n", err)
					client.Close()
					delete(Clients, client)
				}
			}

			InsertMessage(message)
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