package main

import (
	"context"
	"fmt"
	"go-simple-rest/db"
	"go-simple-rest/src"
	"go-simple-rest/src/swagger"
	"go-simple-rest/src/v1/kafkaclient"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	swaggerFiles "github.com/swaggo/files" // swagger embed files
)

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
	initializeDatabase()
	swagger.Initialize()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	src.Routes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// cc, err := natsclient.Consumer()
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer cc.Stop()
	//
	go kafkaclient.ConsumerWithAutoCommit("test-group")
	log.Println("Starting server... on port ", 8080)
	r.Run("127.0.0.1:8080")

}

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

func initializeDatabase() {
	db.Connect()
}
