package transport

import (
	"go-simple-rest/src/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Endpoints struct {
	Login   gin.HandlerFunc
	NewUser gin.HandlerFunc
}

func MakeEndpoints(s auth.Service) Endpoints {
	return Endpoints{
		Login:   makeCreateLoginEndpoint(s),
		NewUser: makeCreateRegisterEndpoint(s),
	}
}

func makeCreateLoginEndpoint(s auth.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		decodedReq, err := decodeCreateLoginRequest(ctx)
		req := decodedReq.(CreateLoginRequest)
		msg, error := s.Login(ctx, req.PAYLOAD)

		if _, ok := error["err"]; ok {
			ctx.IndentedJSON(http.StatusBadRequest, error)
			return
		}

		ctx.SetCookie("token", msg, 3600, "/", "127.0.0.1", false, true)
		ctx.SetSameSite(http.SameSiteStrictMode)
		ctx.JSON(http.StatusOK, CreateLoginResponse{MESSAGE: msg, ERROR: err})
	}
}

func makeCreateRegisterEndpoint(s auth.Service) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		decodedReq, err := decodeCreateRegisterRequest(ctx)
		req := decodedReq.(CreateRegisterRequest)
		msg, error := s.Register(ctx, req.PAYLOAD)

		if _, ok := error["err"]; ok {
			ctx.IndentedJSON(http.StatusBadRequest, error)
			return
		}

		ctx.JSON(http.StatusCreated, CreateRegisterResponse{MESSAGE: msg, ERROR: err})

	}
}

func decodeCreateLoginRequest(c *gin.Context) (request interface{}, err error) {
	var req CreateLoginRequest
	if err := c.BindJSON(&req.PAYLOAD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	return req, nil
}

func decodeCreateRegisterRequest(c *gin.Context) (request interface{}, err error) {
	var req CreateRegisterRequest
	if err := c.BindJSON(&req.PAYLOAD); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	return req, nil
}
