package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rakshit.dev/m/src/configs"
)

var MongoClient *mongo.Database

// initialize mongo client
func init() {
	clientOptions := options.Client().ApplyURI(configs.MongoURI) // Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	MongoClient = client.Database(configs.DBName)
}
