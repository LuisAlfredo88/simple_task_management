package api

import (
	security "stm/modules/security/api"
	taskManagement "stm/modules/task_management/api"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		security.NewSecurityApi,
		taskManagement.NewTaskManagementApi,
	),
)

func RegisterRoutes(
	s *API,
	e *echo.Echo,
	securityApi *security.SecurityApi,
	taskManagementApi *taskManagement.TaskManagementApi,
) {
	go s.Start(e)
	securityApi.Register()
	taskManagementApi.Register()
}
