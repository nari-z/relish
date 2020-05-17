package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/nari-z/relish/authenticator/controller"
	"github.com/nari-z/relish/authenticator/protocol"
)

func main() {
	fmt.Println("Hello gRPC.")

	// launch server
	listenPort, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	service := controller.NewLoginController()
	protocol.RegisterLoginControllerServer(server, service)
	// Register reflection service on gRPC server.
	reflection.Register(server)
	server.Serve(listenPort)
}
