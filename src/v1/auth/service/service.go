package service

import (
	"go-simple-rest/src/v1/auth/model"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Login(ctx *gin.Context, payload model.LoginAuth) (model.AuthResponse, bool)
	Register(ctx *gin.Context, payload model.RegisterAuth) (model.AuthResponse, bool)
}
