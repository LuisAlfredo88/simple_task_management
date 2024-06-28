package service

import (
	securityService "stm/modules/security/service"
	TaskManagementService "stm/modules/task_management/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		// Security services
		securityService.NewSecurityService,
		securityService.NewUserService,

		// Task management services
		TaskManagementService.NewTaskService,
	),
)
