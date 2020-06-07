package controller

import (
	"github.com/nari-z/relish/gateway/generate/restapi/operations"

	"github.com/nari-z/relish/gateway/repository"
)

// Controller bundler controller
type Controller interface {
	SetRoute(api *operations.RelishAPI) *operations.RelishAPI
}

type controller struct {
	healthController HealthController
	loginController  LoginController
}

// NewController create Controller
func NewController() Controller {
	return &controller{
		healthController: NewHealthContoller(),
		loginController:  NewLoginContoller(repository.NewUserRepository()),
	}
}

func (c controller) SetRoute(api *operations.RelishAPI) *operations.RelishAPI {
	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(c.healthController.CheckHealth)
	api.LoginUserHandler = operations.LoginUserHandlerFunc(c.loginController.LoginUser)

	return api
}
