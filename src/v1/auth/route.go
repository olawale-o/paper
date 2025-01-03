package auth

import (
	"go-simple-rest/src/v1/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.Use(middlewares.Logger())
	r.POST("/auth/login", Login)
	r.POST("/auth/sign-up", Register)
}
