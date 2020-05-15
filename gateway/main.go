package main

import (
	"log"

	"github.com/go-openapi/loads"

	"github.com/nari-z/relish/gateway/controller"

	"github.com/nari-z/relish/gateway/generate/restapi"
	"github.com/nari-z/relish/gateway/generate/restapi/operations"
)

func main()  {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewRelishAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.Port = 1111
	defer server.Shutdown()

	controller := controller.NewController()
	api = controller.SetRoute(api)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}