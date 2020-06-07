package controller

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/nari-z/relish/gateway/generate/restapi/operations"
)

// HealthController swagger endpoint
type HealthController interface {
	CheckHealth(operations.CheckHealthParams) middleware.Responder
}

type healthController struct {
}

// NewHealthContoller create HealthController
func NewHealthContoller() HealthController {
	return &healthController{}
}

func (h healthController) CheckHealth(params operations.CheckHealthParams) middleware.Responder {
	fmt.Println("call CheckHealth")
	return operations.NewCheckHealthOK()
}
