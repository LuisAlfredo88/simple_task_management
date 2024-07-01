package contract

import (
	securityEntity "stm/modules/security/domain/entity"
	sharedModel "stm/shared/model"
)

type UserRepository interface {
	Save(user *securityEntity.User) (securityEntity.User, error)
	GetAllUsers(filter *sharedModel.CriteriaFilter) ([]securityEntity.User, int64, error)
	ListUsers() ([]securityEntity.User, error)
	GetUserById(userId string) (securityEntity.User, error)
	ToggleUserStatus(userId string, isActive int8) error
	UserExists(userName string) (bool, error)
}

type UserService interface {
	UserRepository
}
