package route

import (
	"go-simple-rest/src/v1/auth/controller"
	"go-simple-rest/src/v1/auth/model"
	"go-simple-rest/src/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	// r.Use(middlewares.Logger())

	r.POST("/auth/login", middlewares.RequestToJSON[model.LoginAuth](), middlewares.Validator[model.LoginAuth](), controller.Login)

	r.POST("/auth/sign-up", middlewares.RequestToJSON[model.RegisterAuth](), middlewares.Validator[model.RegisterAuth](), controller.Register)

}
