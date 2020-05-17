package controller

import (
	"context"
	"log"

	"github.com/nari-z/relish/authenticator/protocol"
)

type loginController struct {
}

// NewLoginController is create LoginController
func NewLoginController() protocol.LoginControllerServer {
	return &loginController{}
}

func (c loginController) Login(ctx context.Context, in *protocol.LoginRequest) (*protocol.LoginResponse, error) {
	log.Println("call Login")
	return &protocol.LoginResponse{IsOK: true}, nil
}
