package models

import (
	"time"

	"github.com/google/uuid"
)

func NewRoom(id, name string, admin *Agent) *Room {
	return &Room{
		ID:    uuid.New(),
		Name:  name,
		Admin: admin,
		CreatedAt: time.Now(),
		Users: make([]*Agent, 2),
	}
}

type Room struct {
	ID        uuid.UUID
	Name      string
	Admin     *Agent
	CreatedAt time.Time
	Users     []*Agent
}
