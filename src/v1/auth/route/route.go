package route

import (
	"go-simple-rest/src/v1/auth/controller"
	"go-simple-rest/src/v1/auth/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	// r.Use(middlewares.Logger())

	r.Use(middlewares.RequestToJSON(), middlewares.Validator())
	{
		r.POST("/auth/login", controller.Login)
	}
	r.POST("/auth/sign-up", controller.Register)
}
