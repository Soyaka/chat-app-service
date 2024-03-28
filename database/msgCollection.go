package database

import (
	"context"
	"main/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"

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

func GetMessages(filter bson.M) ([]*models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := MsgCollection()

	var messages []*models.Message
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var msg models.Message
		if err := cursor.Decode(&msg); err != nil {
			return nil, err
		}
		messages = append(messages, &msg)
	}
	return messages, nil
}
