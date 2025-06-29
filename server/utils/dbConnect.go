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

	// send db connect request to mongodb
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig.DBURI))
	if err != nil {
		fmt.Println(err)
	}

	// send test db connect request
	if err = client.Ping(ctx, nil); err != nil {
		fmt.Println(err)
	}

	MongoClient = client
	fmt.Println("MONGO DB CONNECTED✅")
	return nil

}