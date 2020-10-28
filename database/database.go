package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Karitham/WaifuBot/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// Init is used to start the database
func Init(config *config.ConfStruct) {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(config.MongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("waifubot").Collection("waifus")
	fmt.Println("Connected to WaifuDB!")
}
