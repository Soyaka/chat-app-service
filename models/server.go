package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	ConnectedUsers map[Agent]*websocket.Conn
	Mu             sync.Mutex
}

func CreateServer() *Server {
	return &Server{
		ConnectedUsers: make(map[Agent]*websocket.Conn),
	}

}
