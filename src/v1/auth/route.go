package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.Engine) {
	r.POST("/auth/login", Login)
	r.POST("/auth/sign-up", Register)
}
