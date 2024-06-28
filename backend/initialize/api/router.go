package api

import (
	"stm/modules/security/api"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(api.NewSecurityApi),
)

func RegisterRoutes(
	s *API,
	e *echo.Echo,
	securityApi *api.SecurityApi,
) {
	go s.Start(e)
	securityApi.Register()
}
