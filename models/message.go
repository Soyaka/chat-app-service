package models

import (
	"time"

	"github.com/google/uuid"
)

func NewMessage(content, from, to, typee string) *Message {
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
	From      string
	To        string
	Type      string
	CreatedAt time.Time
}
