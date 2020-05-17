package controller

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/nari-z/relish/gateway/generate/restapi/operations"
	"google.golang.org/grpc"

	"github.com/nari-z/relish/authenticator/protocol"
)

// LoginController swagger endpoint
type LoginController interface {
	LoginUser(operations.LoginUserParams) middleware.Responder
}

type loginController struct {
}

// NewLoginContoller create LoginController
func NewLoginContoller() LoginController {
	return &loginController{}
}

func (h loginController) LoginUser(params operations.LoginUserParams) middleware.Responder {
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := protocol.NewLoginControllerClient(conn)
	res, err := client.Login(
		params.HTTPRequest.Context(),
		&protocol.LoginRequest{
			Id:       "user id",
			Password: "user password",
		},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	if res.IsOK {
		log.Println("login is OK")
	}
	return operations.NewLoginUserOK()
}
