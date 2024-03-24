package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func NewUser(conn *websocket.Conn) *Agent {
	return &Agent{
		ID:       uuid.New(),
		ChatConn: conn,
	}
}

type Agent struct {
	ID       uuid.UUID
	ChatConn *websocket.Conn
	LastSeen time.Time
}
