package api

import (
	securityService "stm/modules/security/domain/contract"

	"github.com/labstack/echo/v4"
)

type SecurityApi struct {
	securityService securityService.SecurityService
	userService     securityService.UserService
	api             *echo.Echo
}

func NewSecurityApi(
	securityService securityService.SecurityService,
	userService securityService.UserService,
	api *echo.Echo,
) *SecurityApi {
	rest_api := &SecurityApi{
		securityService: securityService,
		userService:     userService,
		api:             api,
	}

	return rest_api
}

func (l *SecurityApi) Register() {
	security := l.api.Group("/security")
	//Registering login handlers
	l.registerLoginHandlers(security)
	//Registering user handlers
	l.registerUserHandlers(security)
}
