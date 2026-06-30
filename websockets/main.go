package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		log.Printf("received: %s", msg)

		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
func main() {
	http.HandleFunc("/", wsHandler)
	http.ListenAndServe(":8080", nil)
}
