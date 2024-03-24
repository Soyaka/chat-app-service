package models

import (
	"time"

	"github.com/google/uuid"
)


func NewConversation( members []*Agent) *Conversation{
	return &Conversation{
		ID : uuid.New(),
		Members: members,
	}
}

type Conversation struct {
	ID uuid.UUID
	Members []*Agent
	CreatedAt time.Time
	LastMessage Message
}
