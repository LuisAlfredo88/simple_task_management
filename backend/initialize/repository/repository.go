package repository

import (
	securityRepo "stm/modules/security/repository"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		securityRepo.NewSecurityRepo,
		securityRepo.NewUserRepo,
	),
)
