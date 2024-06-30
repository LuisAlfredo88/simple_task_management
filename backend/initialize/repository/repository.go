package repository

import (
	securityRepo "stm/modules/security/repository"
	taskManagementRepo "stm/modules/task_management/repository"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		// Security repo
		securityRepo.NewSecurityRepo,
		securityRepo.NewUserRepo,

		// Task managament repo
		taskManagementRepo.NewTaskRepo,
	),
)
