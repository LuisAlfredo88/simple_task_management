package contract

import (
	taskEntity "stm/modules/task_management/domain/entity"
	sharedModel "stm/shared/model"
)

type TaskRepository interface {
	Save(task *taskEntity.Task) (taskEntity.Task, error)
	GetAllTasks(filter *sharedModel.CriteriaFilter) ([]taskEntity.Task, int64, error)
	GetTaskById(taskId uint) (taskEntity.Task, error)
	ChangeTaskStatus(taskId uint, status uint) error
}

type TaskService interface {
	TaskRepository
}
