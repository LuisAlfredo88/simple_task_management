package repository

import (
	securityService "stm/modules/security/domain/contract"
	securityEntity "stm/modules/security/domain/entity"
	sharedModel "stm/shared/model"
	encoding "stm/shared/utility/encoding"
	gormUtil "stm/shared/utility/gorm"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func (l *userRepo) Save(user *securityEntity.User) (securityEntity.User, error) {
	if user.Id == "" {
		user.Id = uuid.New().String()
		user.Password = encoding.GetMD5Hash(user.Password)
	}

	err := l.db.Where(securityEntity.User{Id: user.Id}).
		Assign(user).
		FirstOrCreate(&securityEntity.User{}).Error

	return *user, err
}

func (l *userRepo) GetAllUsers(filter *sharedModel.CriteriaFilter) ([]securityEntity.User, int64, error) {
	users := []securityEntity.User{}

	query := l.db.Model(securityEntity.User{})

	if search, ok := (filter.Filters)["search"]; ok && search != "" {
		query = query.Where("concat(name, ' ', user_name, ' ', last_name) LIKE ?", "%"+search.(string)+"%")
	}

	if isActive, ok := (filter.Filters)["isActive"]; ok && isActive != "" {
		query = query.Where("is_active = ?", isActive)
	}

	// Getting records count
	total := gormUtil.GetCount(l.db, query)

	query = query.Limit(int(filter.Limit))
	query = query.Offset(int(filter.Skip))

	query = query.Order("name desc")
	l.db.Raw("?", query).Scan(&users)

	return users, total, nil
}

func (l *userRepo) GetUserById(userId string) (securityEntity.User, error) {
	user := securityEntity.User{}

	err := l.db.Where(&securityEntity.User{
		Id: userId,
	}).Find(&user).Error

	user.Password = ""

	return user, err
}

func (l *userRepo) ToggleUserStatus(userId string, isActive int8) error {
	user := securityEntity.User{}

	err := l.db.Model(&user).
		Where("id = ?", userId).
		Update("is_active", isActive).Error

	return err
}

func (l *userRepo) UserExists(userName string) (bool, error) {
	var exists bool
	err := l.db.Model(&securityEntity.User{}).
		Select("count(*) > 0").
		Where("user_name = ?", userName).
		Find(&exists).
		Error

	return exists, err
}

func (l *userRepo) ListUsers() ([]securityEntity.User, error) {
	var users = []securityEntity.User{}
	err := l.db.Model(&securityEntity.User{}).
		Find(&users).
		Error

	return users, err
}

func NewUserRepo(db *gorm.DB) securityService.UserRepository {
	return &userRepo{
		db: db,
	}
}
