package transport

import "go-simple-rest/src/v1/auth"

type CreateLoginRequest struct {
	PAYLOAD auth.LoginAuth
}

type CreateLoginResponse struct {
	MESSAGE string
	ERROR   error
}

type CreateRegisterRequest struct {
	PAYLOAD auth.RegisterAuth
}

type CreateRegisterResponse struct {
	MESSAGE string
	ERROR   error
}
