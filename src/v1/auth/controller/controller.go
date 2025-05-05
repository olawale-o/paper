package controller

import "github.com/gin-gonic/gin"

// Login godoc
// @Tags User Authentication
// @Summary Login user
// @Description Login user with username and password
// @Accept json
// @Param data body model.LoginAuth true "User"
// @Produce json
// @Success 200 {object} model.LoginResponse "Response"
// @Header 200 {string} Cookie "session_id"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /auth/login [post]
type Controller interface {
	Login(c *gin.Context)

	// Sign up godoc
	// @Tags User Authentication
	// @Summary Register user
	// @Description Register user with username,password, firstname and lastname
	// @Accept json
	// @Param data body model.RegisterAuth true "User"
	// @Produce json
	// @Success 201 {object} string "Response"
	// @Failure 400 {object} string "Error"
	// @Failure 500 {object} string "Error"
	// @Router /auth/sign-up [post]
	Register(c *gin.Context)
}
