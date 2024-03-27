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
	Id        uuid.UUID `json:"-" bson:"_id"`
	Content   string    `json:"content" bson:"content"`
	From      string    `json:"from" bson:"from"`
	To        string    `json:"to" bson:"to"`
	Type      string    `json:"-" bson:"type"`
	CreatedAt time.Time `json:"-" bson:"created_at"`
}
