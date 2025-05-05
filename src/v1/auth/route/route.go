package route

import (
	"go-simple-rest/src/v1/auth/controller"
	"go-simple-rest/src/v1/auth/model"
	"go-simple-rest/src/v1/middlewares"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthRoutes(r *gin.RouterGroup, databaseClient *mongo.Database) {
	// r.Use(middlewares.Logger())
	//
	controller := controller.AuthControllerImpl(databaseClient)

	r.POST("/auth/login", middlewares.RequestToJSON[model.LoginAuth](), middlewares.Validator[model.LoginAuth](), controller.Login)

	r.POST("/auth/sign-up", middlewares.RequestToJSON[model.RegisterAuth](), middlewares.Validator[model.RegisterAuth](), controller.Register)

}
