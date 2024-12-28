package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

func Connect() (*mongo.Client, context.Context, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetServerAPIOptions(serverAPI)
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		fmt.Println("It is callled here")
		panic(err)
	}

	var result bson.M
	if err := client.Database("admin").RunCommand(ctx, bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client, ctx, err
}
