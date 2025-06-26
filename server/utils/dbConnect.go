package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/AbdulRahman-04/Go_Backend_Practice/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var MongoClient *mongo.Client

func DbConnect() error {
	// create ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create connection with mongo db 
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig.DBURI))
	if err != nil {
		fmt.Println(err)
	}

	// send test request to db 
	if err = client.Ping(ctx, nil); err != nil {
		fmt.Println(err)
	}

	// assign client to mongoClient
	MongoClient = client

	fmt.Println("DB CONNECTED✅")
	return nil
}