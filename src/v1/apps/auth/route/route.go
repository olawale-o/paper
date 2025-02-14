package route

import (
	"auth/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	// r.Use(middlewares.Logger())
	//
	r.POST("/login", controller.Login)
	r.POST("/sign-up", controller.Register)

}
