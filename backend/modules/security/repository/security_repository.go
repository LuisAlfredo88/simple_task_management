package repository

import (
	securityService "stm/modules/security/domain/contract"
	dto "stm/modules/security/domain/dto"
	entity "stm/modules/security/domain/entity"
	encoding "stm/shared/utility/encoding"

	"gorm.io/gorm"
)

type securityRepo struct {
	db *gorm.DB
}

func (l *securityRepo) Authenticate(loginData dto.LoginData) (entity.User, error) {
	password := encoding.GetMD5Hash(loginData.Password)
	user := entity.User{}

	err := l.db.Where(entity.User{
		UserName: loginData.UserName,
		Password: password,
		IsActive: 1,
	}).First(&user).Error

	return user, err
}

func NewSecurityRepo(db *gorm.DB) securityService.SecurityRepository {
	return &securityRepo{
		db: db,
	}
}
