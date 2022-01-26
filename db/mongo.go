package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
const uri = "mongodb://127.0.0.1:27017"

type DB struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func ConnectDB() *mongo.Client {
	fmt.Println("Connecting to mongodb")
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	return client
}
