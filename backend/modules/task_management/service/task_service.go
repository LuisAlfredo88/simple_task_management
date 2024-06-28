package service

import (
	taskManagementContract "stm/modules/task_management/domain/contract"
	taskManagementEntity "stm/modules/task_management/domain/entity"
	sharedModel "stm/shared/model"
)

type task struct {
	taskRepo           taskManagementContract.TaskRepository
	taskManagementRepo taskManagementContract.TaskRepository
}

func (l *task) Save(task *taskManagementEntity.Task) (taskManagementEntity.Task, error) {
	if err := task.Validate(); err != nil {
		return taskManagementEntity.Task{}, err
	}

	return l.taskRepo.Save(task)
}

func (l *task) GetAllTasks(filter *sharedModel.CriteriaFilter) ([]taskManagementEntity.Task, int64, error) {
	return l.taskRepo.GetAllTasks(filter)
}

func (l *task) GetTaskById(taskId uint) (taskManagementEntity.Task, error) {
	task, err := l.taskRepo.GetTaskById(taskId)

	if err != nil {
		return taskManagementEntity.Task{}, err
	}

	return task, err
}

func (l *task) ChangeTaskStatus(taskId uint, status uint) error {
	return l.taskRepo.ChangeTaskStatus(taskId, status)
}

func NewTaskService(
	taskRepo taskManagementContract.TaskRepository,
	taskManagementRepo taskManagementContract.TaskRepository,
) taskManagementContract.TaskService {
	return &task{
		taskRepo:           taskRepo,
		taskManagementRepo: taskManagementRepo,
	}
}
