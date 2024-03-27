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

