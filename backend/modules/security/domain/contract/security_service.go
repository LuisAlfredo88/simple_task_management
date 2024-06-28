package contract

import (
	securityDto "stm/modules/security/domain/dto"
	securityEntity "stm/modules/security/domain/entity"
)

type SecurityRepository interface {
	Authenticate(loginInfo securityDto.LoginData) (securityEntity.User, error)
}

type SecurityService interface {
	Authenticate(loginInfo securityDto.LoginData) (*securityDto.LoginResponseData, error)
	RefreshToken(userId string) (string, error)
}
