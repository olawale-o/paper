package main

import (
	"fmt"
	"go-simple-rest/db"
	"go-simple-rest/src"
	"log"

	"github.com/gin-gonic/gin"
)

func terminate() {
	if r := recover(); r != nil {
		fmt.Println("An error occured ", r)
		fmt.Println("Application terminated gracefully")
	} else {
		fmt.Println("Application executed succcesfully")
	}
}
func main() {
	defer terminate()
	client, ctx, err := db.Connect()
	// defer cancel()
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	r := gin.Default()
	src.Routes(r)
	log.Println("Starting server... on port ", 8080)
	r.Run("localhost:8080")
}
