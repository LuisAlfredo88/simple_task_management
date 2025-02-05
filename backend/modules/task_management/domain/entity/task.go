package domain

import (
	"errors"
	securityEntity "stm/modules/security/domain/entity"
)

type TaskStatus struct {
	Id    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"size:100" json:"name"`
	Color string `gorm:"size:7" json:"color"`
}

func (TaskStatus) TableName() string {
	return "task_status"
}

type Task struct {
	Id           uint                `gorm:"primaryKey; autoIncrement" json:"id"`
	Title        string              `gorm:"size:100" json:"title"`
	Description  string              `gorm:"size:255" json:"description"`
	Status       TaskStatus          `json:"status"`
	StatusId     uint                `json:"statusId"`
	CreatedBy    securityEntity.User `json:"-"`
	CreatedById  string              `json:"createdBy"`
	AssignedTo   securityEntity.User `json:"-"`
	AssignedToId string              `gorm:"default:null" json:"assignedTo"`
}

func (u *Task) Validate() error {
	if u.Title == "" {
		return errors.New("must specify task title")
	}

	if u.Description == "" {
		return errors.New("must specify task description")
	}

	if u.StatusId == 0 && u.Status.Id == 0 {
		return errors.New("must specify status")
	}

	if u.Id == 0 && u.CreatedById == "" && u.CreatedBy.Id == "" {
		return errors.New("must specify created by")
	}

	return nil
}
