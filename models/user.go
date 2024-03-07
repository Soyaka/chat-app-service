package models

import "github.com/google/uuid"

func NewUser(id, name string) *User {
	return &User{
		ID:    uuid.New(),
		Name:  name,
		Rooms: make(map[string]*Room),
	}
}

type User struct {
	ID    uuid.UUID
	Name  string
	Rooms map[string]*Room
}


