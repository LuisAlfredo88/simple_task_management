package contract

import (
	taskDto "stm/modules/task_management/domain/dto"
	taskEntity "stm/modules/task_management/domain/entity"
	sharedModel "stm/shared/model"
)

type TaskRepository interface {
	Save(task *taskEntity.Task) (taskEntity.Task, error)
	GetAllTasks(filter *sharedModel.CriteriaFilter) ([]taskDto.TaskRecord, int64, error)
	GetTaskById(taskId uint) (taskEntity.Task, error)
	ChangeTaskStatus(taskId uint, status uint) error
	GetAllTasksStatus() ([]taskEntity.TaskStatus, error)
}

type TaskService interface {
	Save(task *taskEntity.Task, userId string) (taskEntity.Task, error)
	GetAllTasks(filter *sharedModel.CriteriaFilter) ([]taskDto.TaskRecord, int64, error)
	GetTaskById(taskId uint) (taskEntity.Task, error)
	ChangeTaskStatus(taskId uint, status uint) error
	GetAllTasksStatus() ([]taskEntity.TaskStatus, error)
}
