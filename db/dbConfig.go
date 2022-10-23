package db

import (
	"context"
	"fmt"
	"log"

	"github.com/vimalkumar-2124/sample-authentication/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create mongoDb connection
func CreateDbConnection() (*mongo.Client, error) {
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.lowctzs.mongodb.net/?retryWrites=true&w=majority", config.EnvConfig("DB_USER"), config.EnvConfig("DB_PASS")))
	// Connect to MongoDb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDb")
	return client, err
}
