package database

import (
	"context"
	"main/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

//TODO: add the cursor limit

func GetUsers(filter bson.M) ([]*models.Agent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := UserCollection()
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var users []*models.Agent
	for cursor.Next(ctx) {
		var user *models.Agent
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
