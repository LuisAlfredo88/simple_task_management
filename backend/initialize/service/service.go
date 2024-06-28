package service

import (
	securityService "stm/modules/security/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		securityService.NewSecurityService,
		securityService.NewUserService,
	),
)
