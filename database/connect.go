package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		panic(err)
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB_URL")))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client

}



