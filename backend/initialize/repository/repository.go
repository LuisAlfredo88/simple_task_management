package repository

import (
	securityRepo "stm/modules/security/repository"
	taskManagemtRepo "stm/modules/task_management/repository"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		// Security repo
		securityRepo.NewSecurityRepo,
		securityRepo.NewUserRepo,

		// Task managament repo
		taskManagemtRepo.NewTaskRepo,
	),
)
