package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/auth/login", Login)
	r.POST("/auth/sign-up", Register)
}
