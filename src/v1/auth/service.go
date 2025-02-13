package auth

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	Login(ctx *gin.Context, auth LoginAuth) (string, gin.H)
	Register(ctx *gin.Context, auth RegisterAuth) (string, gin.H)
}
