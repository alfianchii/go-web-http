package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

var (
	Clients = make(map[*websocket.Conn]string)
	Typers = make(map[string]bool)
	ClientsLock sync.Mutex
	// This is a ref just like on Nuxt
	BroadcastChan = make(chan ClientChat)
)