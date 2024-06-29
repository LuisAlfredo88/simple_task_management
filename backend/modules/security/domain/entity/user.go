package domain

import (
	"encoding/json"
	"errors"
)

type User struct {
	Id       string `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"size:255" json:"name"`
	LastName string `gorm:"size:255" json:"lastName"`
	UserName string `gorm:"size:50" json:"userName"`
	Password string `gorm:"size:255" json:"password,omitempty"`
	// IsActive field usually goes boolean, but bool is not supported on Sqlite
	IsActive int8 `json:"isActive"`
}

func (p User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Alias
		Password string `json:"password,omitempty"` // Excluding Password from JSON
	}{
		Alias:    (Alias)(p),
		Password: "",
	})
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("must specify name")
	}

	if u.UserName == "" {
		return errors.New("must specify user name")
	}

	return nil
}
