package models

import (
	"errors"

	"github.com/google/uuid"
)

type Agent struct {
	ID       uuid.UUID `json:"-"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func CreateAgent(username, email  string) (*Agent, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	agent := &Agent{
		Username: username,
		Email:    email,


	}
	return agent, nil
}
