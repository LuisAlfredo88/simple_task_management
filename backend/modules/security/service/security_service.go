package service

import (
	"errors"
	securityContract "stm/modules/security/domain/contract"
	dto "stm/modules/security/domain/dto"
	jwt "stm/shared/utility/jwt"
)

type login struct {
	securityRepo securityContract.SecurityRepository
}

func getToken(userId string) (string, error) {
	// Building JWT
	jwtData := map[string]interface{}{
		"userId": userId,
	}

	return jwt.GenerateJWT(jwtData)
}

func (l *login) Authenticate(loginData dto.LoginData) (*dto.LoginResponseData, error) {
	user, err := l.securityRepo.Authenticate(loginData)

	if err != nil {
		return nil, errors.New("login error")
	}

	token, err := getToken(user.Id)

	if err != nil {
		return nil, err
	}

	// Getting logging information
	loginInfo := &dto.LoginResponseData{
		User: dto.ExposedUserData{
			Id:       user.Id,
			UserName: user.UserName,
		},
		Token: token,
	}

	return loginInfo, nil
}

func (l *login) RefreshToken(userId string) (string, error) {
	token, err := getToken(userId)

	if err != nil {
		return "", err
	}

	return token, nil
}

func NewSecurityService(securityRepo securityContract.SecurityRepository) securityContract.SecurityService {
	return &login{
		securityRepo: securityRepo,
	}
}
