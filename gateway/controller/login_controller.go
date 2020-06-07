package controller

import (
	"fmt"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/nari-z/relish/gateway/generate/restapi/operations"

	"github.com/nari-z/relish/gateway/domain"
)

// LoginController swagger endpoint
type LoginController interface {
	LoginUser(operations.LoginUserParams) middleware.Responder
}

type loginController struct {
	userRepository domain.UserRepository
}

// NewLoginContoller create LoginController
func NewLoginContoller(userRepository domain.UserRepository) LoginController {
	return &loginController{userRepository}
}

func (c loginController) LoginUser(params operations.LoginUserParams) middleware.Responder {
	fmt.Println("call CheckHealth")

	isOK, err := c.userRepository.Login("user id", "user password", params.HTTPRequest.Context())

	if err != nil {
		log.Fatal(err.Error())
	}
	if isOK {
		log.Println("login is OK")
	}
	return operations.NewLoginUserOK()
}
