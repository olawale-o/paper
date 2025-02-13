package route

import (
	"go-simple-rest/src/v1/auth/transport"

	"github.com/go-kit/log"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, svcEndpoint transport.Endpoints, logger log.Logger) {
	// r.Use(middlewares.Logger())
	r.POST("/auth/login", svcEndpoint.Login)
	r.POST("/auth/sign-up", svcEndpoint.NewUser)
}
