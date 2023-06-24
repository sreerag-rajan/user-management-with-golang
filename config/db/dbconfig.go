package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//holds the configuration details
type MongoDBConfig struct {
	ConnectionUrl string
}

//holds the mongo connection client
type MongoDBConnection struct {
	Client              *mongo.Client
	isConnectionSuccess bool
}

func NewMongoDBConnection(config *MongoDBConfig) (*MongoDBConnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(config.ConnectionUrl)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	return &MongoDBConnection{
		Client:              client,
		isConnectionSuccess: true,
	}, nil
}

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connectionUrl := os.Getenv("DB_URL")

	if connectionUrl == "" {
		log.Fatal("NO DB URL")
	}

	connConfig := &MongoDBConfig{
		ConnectionUrl: connectionUrl,
	}

	mongoConnection, err := NewMongoDBConnection(connConfig)

	if err != nil {
		log.Fatal("Mongo Connection Failed")
	}

	if mongoConnection.isConnectionSuccess {
		fmt.Println("MONGO DB SUCCESSFULLY CONNECTED")
	}
}
