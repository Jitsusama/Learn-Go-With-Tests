package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type incomingSocket struct {
	*websocket.Conn
}

func newIncomingSocket(w http.ResponseWriter, r *http.Request) *incomingSocket {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrading to WebSockets failed: %v\n", err)
	}
	return &incomingSocket{conn}
}

func (s *incomingSocket) WaitForMessage() string {
	_, msg, err := s.ReadMessage()
	if err != nil {
		log.Printf("WebSocket read failed: %v\n", err)
	}
	return string(msg)
}

func (s *incomingSocket) Write(p []byte) (n int, err error) {
	err = s.WriteMessage(websocket.TextMessage, p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
