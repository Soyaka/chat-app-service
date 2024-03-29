package database

import (
	"context"
	"main/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func UserCollection(client *mongo.Client) *mongo.Collection {

	coll := client.Database("chat-app").Collection("users")
	return coll
}

func InsertUser(user *models.User , client *mongo.Client) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	coll := UserCollection(client)
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//TODO: add the cursor limit

func GetUsers(filter bson.M, client *mongo.Client) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := UserCollection(client)
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var users []*models.User
	for cursor.Next(ctx) {
		var user *models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUser(filter bson.M , Client *mongo.Client) (*models.User, error) {
	var users []*models.User
	users, err := GetUsers(filter , Client)
	if err != nil {
		return nil, err
	}
	return users[0], nil
}
