package models

import (
	"time"

	"github.com/google/uuid"
)

func NewMessage(content string, from *Agent, to *Agent, typee string) *Message {
	return &Message{
		Id:        uuid.New(),
		Content:   content,
		From:      from,
		To:        to,
		Type:      typee,
		CreatedAt: time.Now(),
	}
}

type Message struct {
	Id        uuid.UUID
	Content   string
	From      *Agent
	To        *Agent
	Type      string
	CreatedAt time.Time
}
