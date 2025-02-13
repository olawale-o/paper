package service

import (
	"go-simple-rest/src/v1/auth"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Login(ctx *gin.Context, payload auth.LoginAuth) (string, gin.H)
	Register(ctx *gin.Context, payload auth.RegisterAuth) (string, gin.H)
}
