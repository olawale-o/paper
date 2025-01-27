package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	// r.Use(middlewares.Logger())
	r.POST("/auth/login", Login)
	r.POST("/auth/sign-up", Register)
}
