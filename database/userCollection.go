package database

import (
	"context"
	"main/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func UserCollection() *mongo.Collection {
	client := Connect()
	coll := client.Database("chat-app").Collection("users")
	return coll
}

func InsertUser(user *models.Agent) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	coll := UserCollection()
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
