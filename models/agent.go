package models

import (
	"errors"

	"github.com/google/uuid"
)

type Agent struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func CreateUser(username, email, password string) (*Agent, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	user := &Agent{
		Username: username,
		Email:    email,
		Password: password,

	}
	return user, nil
}
