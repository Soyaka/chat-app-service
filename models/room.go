package models

import (
	"time"

	"github.com/google/uuid"
)

func NewRoom(id, name string, admin *User) *Room {
	return &Room{
		ID:    uuid.New(),
		Name:  name,
		Admin: admin,
		CreatedAt: time.Now(),
		Users: make(map[string]*User),
	}
}

type Room struct {
	ID        uuid.UUID
	Name      string
	Admin     *User
	CreatedAt time.Time
	Users     map[string]*User
}
