package service

import (
	"auth/model"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Login(ctx *gin.Context, payload model.LoginAuth) (string, gin.H)
	Register(ctx *gin.Context, payload model.RegisterAuth) (string, gin.H)
}
