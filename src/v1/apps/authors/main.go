package main

import (
	"authors/db"
	"authors/kafkaclient"
	"authors/route"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func terminate() {
	if r := recover(); r != nil {
		fmt.Println("An error occured ", r)
		fmt.Println("Application terminated gracefully")
	} else {
		fmt.Println("Application executed succcesfully")
	}
}

func iterateChangeStream(routineCtx context.Context, waitGroup *sync.WaitGroup, stream *mongo.ChangeStream) {
	defer stream.Close(routineCtx)
	defer waitGroup.Done()
	for stream.Next(routineCtx) {
		var data bson.M
		if err := stream.Decode(&data); err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", data)
	}
}

func initializeDB() {
	db.Connect()
}

func main() {
	initializeDB()

	r := gin.New()
	route.AuthorRoutes(r)
	s := &http.Server{
		Addr:           "localhost:8082",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go kafkaclient.ConsumerWithAutoCommit("article-group")
	log.Println("Starting server... on port ", 8082)
	s.ListenAndServe()
}
