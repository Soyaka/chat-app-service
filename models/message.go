package models

import (
	"time"

	"github.com/google/uuid"
)

func NewMessage(content string, from, to string, typee string) *Message {
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
	Id        uuid.UUID `json:"-"`
	Content   string    `json:"content"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Type      string    `json:"-"`
	CreatedAt time.Time `json:"-"`
}
