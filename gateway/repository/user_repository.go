package repository

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/nari-z/relish/gateway/domain"
	"github.com/nari-z/relish/authenticator/protocol"
)

type userRepository struct {
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

func (userRepository) Login(id, password string, ctx context.Context) (bool, error) {
	conn, err := grpc.Dial(os.Getenv("AUTHENTICATOR_HOST"), grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	
	client := protocol.NewLoginControllerClient(conn)
	res, err := client.Login(
		ctx,
		&protocol.LoginRequest{
			Id:       id,
			Password: password,
		},
	)

	return res.IsOK, err
}
