package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type doc map[string]string

type documentKey struct {
	ID primitive.ObjectID `bson:"_id"`
}

type changeID struct {
	Data string `bson:"_data"`
}

type namespace struct {
	Db   string `bson:"db"`
	Coll string `bson:"coll"`
}

type changeEvent struct {
	ID            changeID            `bson:"_id" json:"id"`
	OperationType string              `bson:"operationType" json:"operationType"`
	ClusterTime   primitive.Timestamp `bson:"clusterTime" json:"clusterTime"`
	FullDocument  doc                 `bson:"fullDocument" json:"fullDocument"`
	DocumentKey   documentKey         `bson:"documentKey" json:"documentKey"`
	Ns            namespace           `bson:"ns" json:"ns"`
}

func watch(collection *mongo.Collection) {
	cs, err := collection.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		fmt.Println(err.Error())
	}
	var changeEvent changeEvent
	for cs.Next(context.TODO()) {
		err := cs.Decode(&changeEvent)
		if err != nil {
			log.Fatal(err)
		}

		if changeEvent.OperationType == "insert" {
			fmt.Println(cs.Current)
			// insertArticleToUser(collection, changeEvent.FullDocument)
		}
	}

}

func insertArticleToUser(client *mongo.Client, doc map[string]string) {
	collection := client.Database("go").Collection("users")
	authorOid, _ := primitive.ObjectIDFromHex(doc["authorId"])
	filter := bson.D{{"_id", authorOid}}
	update := bson.D{{"$push", bson.D{{"articles", doc}}}}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	var updatedDoc interface{}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDoc)
	if err != nil {
		panic(err)
	}

	res, _ := bson.MarshalExtJSON(updatedDoc, false, false)
	fmt.Println(string(res))
}

func watchDb(client *mongo.Client, ctx context.Context) {
	// Watch the database
	var changeEvent changeEvent
	cs, err := client.Watch(ctx, mongo.Pipeline{})
	if err != nil {
		fmt.Println(err.Error())
	}

	defer cs.Close(ctx)
	for cs.Next(ctx) {
		err := cs.Decode(&changeEvent)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(cs)

		// if changeEvent.OperationType == "insert" && changeEvent.Ns.Coll == "articles" {
		// 	// insertArticleToUser(client, changeEvent.FullDocument)
		// 	// fmt.Println(changeEvent.FullDocument)
		// 	//
		// 	fmt.Println(changeEvent)
		// }
	}
	// fmt.Println(changeEvent)

}

// const uri = "mongodb://localhost:27017"
const uri = "mongodb://localhost:27018,localhost:27019,localhost:27010/?replicaSet=rs0"

func LoadConfig(uri string) *options.ClientOptions {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	return opts
}

func Connect() (*mongo.Client, context.Context, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
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

	// Get a handle for your collection
	// collection := client.Database("go").Collection("articles")

	// Watches the todo collection and prints out any changed documents
	// go watch(collection)

	return client, ctx, err
}
