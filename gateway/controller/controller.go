package controller

import (
	"github.com/nari-z/relish/gateway/generate/restapi/operations"
)

type Controller interface {
	SetRoute(api *operations.RelishAPI) *operations.RelishAPI
}

type controller struct {
	healthController HealthController
}

func NewController() Controller {
	return &controller{
		healthController: &healthController{},
	}
}

func (c controller) SetRoute(api *operations.RelishAPI) *operations.RelishAPI {
	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(c.healthController.CheckHealth)

	return api
}
