package domain

import (
	"errors"
)

type TaskStatus struct {
	Id   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"size:100" json:"name"`
}

func (TaskStatus) TableName() string {
	return "task_status"
}

type Task struct {
	Id          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string     `gorm:"size:100" json:"name"`
	Description string     `gorm:"size:255" json:"description"`
	Status      TaskStatus `json:"status"`
	StatusId    uint       `json:"status_id"`
}

func (u *Task) Validate() error {
	if u.Title == "" {
		return errors.New("must specify task title")
	}

	if u.Description == "" {
		return errors.New("must specify task description")
	}

	return nil
}
