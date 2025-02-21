package main

import (
	"context"
	"fmt"
	"go-simple-rest/db"
	docs "go-simple-rest/docs"
	"go-simple-rest/src"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	swaggerFiles "github.com/swaggo/files" // swagger embed files
)

func terminate() {
	if r := recover(); r != nil {
		fmt.Println("An error occured ", r)
		fmt.Println("Application terminated gracefully")
	} else {
		fmt.Println("Application executed succcesfully")
	}
}

func iterateChangeStream(routineCtx context.Context, waitGroup sync.WaitGroup, stream *mongo.ChangeStream) {
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

func init() {

	// defer terminate()
	db.Connect()
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		127.0.0.1:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := gin.Default()
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "A simple REST API for managing users and posts."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	src.Routes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Starting server... on port ", 8080)
	r.Run("127.0.0.1:8080")

}
