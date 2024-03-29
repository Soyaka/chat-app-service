package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"-" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"-" bson:"updated_at"`
	DeletedAt time.Time          `json:"-" bson:"deleted_at"`
}

func NewUser(username, email, password string) *User {
	return &User{
		ID:        primitive.NewObjectID(),
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
	}
}
