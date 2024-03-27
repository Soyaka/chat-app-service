package database

import (
	"context"
	"main/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GroupCollection() *mongo.Collection {
	client := Connect()
	coll := client.Database("chat-app").Collection("groups")
	return coll
}

func InsertGroup(group *models.Room) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	coll := GroupCollection()
	res, err := coll.InsertOne(ctx, group)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func GetGroups(filter bson.M) ([]*models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	coll := GroupCollection()
	var group *models.Room
	var groups []*models.Room
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		if err := cursor.Decode(&group); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
