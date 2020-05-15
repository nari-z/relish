package controller

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/nari-z/relish/gateway/generate/restapi/operations"
)

type HealthController interface {
	CheckHealth(operations.CheckHealthParams) middleware.Responder
}

type healthController struct {

}

func (h healthController) CheckHealth(params operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK()
}