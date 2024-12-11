package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

func BroadcastMessages() {
	for {
		msg := <-BroadcastChan

		BroadcastTypers()

		if msg.Send {
			ClientsLock.Lock()
			for client := range Clients {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("Error writing message: %v\n", err)
					client.Close()
					delete(Clients, client)
				}
			}
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