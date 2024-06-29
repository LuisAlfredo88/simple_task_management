package service

import (
	"errors"
	securityContract "stm/modules/security/domain/contract"
	securityEntity "stm/modules/security/domain/entity"
	sharedModel "stm/shared/model"
)

type user struct {
	userRepo     securityContract.UserRepository
	securityRepo securityContract.SecurityRepository
}

func (l *user) Save(user *securityEntity.User) (securityEntity.User, error) {
	if err := user.Validate(); err != nil {
		return securityEntity.User{}, err
	}

	if user.Id == "" {
		user.IsActive = 1

		if exists, _ := l.userRepo.UserExists(user.UserName); exists {
			return securityEntity.User{}, errors.New("user already exists")
		}
	}

	return l.userRepo.Save(user)
}

func (l *user) GetAllUsers(filter *sharedModel.CriteriaFilter) ([]securityEntity.User, int64, error) {
	return l.userRepo.GetAllUsers(filter)
}

func (l *user) GetUserById(userId string) (securityEntity.User, error) {
	user, err := l.userRepo.GetUserById(userId)

	if err != nil {
		return securityEntity.User{}, err
	}

	return user, err
}

func (l *user) ToggleUserStatus(userId string, isActive int8) error {
	return l.userRepo.ToggleUserStatus(userId, isActive)
}

func (l *user) UserExists(userName string) (bool, error) {
	return l.userRepo.UserExists(userName)
}

func (l *user) ListUsers() ([]securityEntity.User, error) {
	return l.userRepo.ListUsers()
}

func NewUserService(
	userRepo securityContract.UserRepository,
	securityRepo securityContract.SecurityRepository,
) securityContract.UserService {
	return &user{
		userRepo:     userRepo,
		securityRepo: securityRepo,
	}
}
