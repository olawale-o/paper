package http

import (
	"context"
	"encoding/json"
	"go-simple-rest/src"
	"go-simple-rest/src/v1/auth/transport"

	"net/http"

	"github.com/go-kit/log"

	"github.com/gin-gonic/gin"
)

func NewService(svcEndpoints transport.Endpoints, logger log.Logger) *gin.Engine {
	r := gin.New()

	src.Routes(r, svcEndpoints, logger)

	return r
}

//
//

//	func decodeCreateLoginRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
//		var req transport.CreateLoginRequest
//		if e := json.NewDecoder(r.Body).Decode(&req.PAYLOAD); e != nil {
//			return nil, e
//		}
//		return req, nil
//	}
// func decodeCreateLoginRequest(c *gin.Context) (request interface{}, err error) {
// 	var req transport.CreateLoginRequest
// 	// if e := json.NewDecoder(r.Body).Decode(&req.PAYLOAD); e != nil {
// 	// 	return nil, e
// 	// }
// 	// return req, nil

// 	if err := c.BindJSON(&req.PAYLOAD); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return nil, err
// 	}
// 	return req, nil
// }

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// func codeFrom(err error) int {
// 	switch err {
// 	case order.ErrOrderNotFound:
// 		return http.StatusBadRequest
// 	default:
// 		return http.StatusInternalServerError
// 	}
// }
