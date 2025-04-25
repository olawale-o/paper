package route

import (
	"go-simple-rest/src/v1/auth/controller"
	"go-simple-rest/src/v1/auth/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	// r.Use(middlewares.Logger())

	r.POST("/auth/login", middlewares.RequestToJSON("login"), middlewares.Validator("login"), controller.Login)

	r.POST("/auth/sign-up", middlewares.RequestToJSON("register"), middlewares.Validator("register"), controller.Register)

}
