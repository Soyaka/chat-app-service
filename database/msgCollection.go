package database

import (
	"context"
	"main/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func MsgCollection() *mongo.Collection {
	client := Connect()
	coll := client.Database("chat-app").Collection("messages")
	return coll
}

func InsertMsg(msg *models.Message) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := MsgCollection().InsertOne(ctx, msg)
	if err != nil {
		return nil, err
	}
	return res, nil

}
