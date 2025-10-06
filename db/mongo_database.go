package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	dbName DBName
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

// ConnectMongo initializes MongoDB client once
func ConnectMongo() *mongo.Client {
	mongoUri := os.Getenv("MONGO_URI")

	if mongoUri == "" {
		mongoUri = "mongodb://localhost:27017"
	}

	clientOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(mongoUri)
		var err error

		//Add timeout context
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		clientInstance, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
		}

		err = clientInstance.Ping(ctx, nil)
		if err != nil {
			log.Fatal("Failed to ping MongoDB:", err)
		}

		fmt.Println("MongoDB client initialized!")
	})

	return clientInstance
}

func ConnectToDB(db DBName) {
	Client = ConnectMongo()
	dbName = db
}

func GetCollection(collection string) *mongo.Collection {
	return Client.Database(string(dbName)).Collection(collection)
}
