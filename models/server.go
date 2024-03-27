package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	ConnectedUsers map[*Agent]*websocket.Conn
	Mu             sync.Mutex
}

func CreateServer() *Server {
	return &Server{
		ConnectedUsers: make(map[*Agent]*websocket.Conn),
	}
}

func (s *Server) AddUser(user *Agent, conn *websocket.Conn) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.ConnectedUsers[user] = conn
}

func (s *Server) RemoveUser(user *Agent) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.ConnectedUsers, user)
}

func (S *Server) GetConnectedUsers() []*Agent {
	var ListUsers []*Agent
	for user := range S.ConnectedUsers {
		ListUsers = append(ListUsers, user)
	}
	return ListUsers
}
