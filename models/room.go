package models

import "github.com/google/uuid"

func NewRoom(id, name string) *Room {
	return &Room{
		ID:    uuid.New(),
		Name:  name,
		Users: make(map[string]*User),
	}
}

type Room struct {
	ID uuid.UUID

	Name  string
	Users map[string]*User
}
