package models

import (
	"github.com/google/uuid"
)

type Agent struct {
	ID       uuid.UUID `json:"-"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}
