package repository

import (
	"errors"
	taskManagementContract "stm/modules/task_management/domain/contract"
	taskDto "stm/modules/task_management/domain/dto"
	taskManagementEntity "stm/modules/task_management/domain/entity"
	sharedModel "stm/shared/model"
)

type taskRepo struct{}

func (l *taskRepo) Save(task *taskManagementEntity.Task) (taskManagementEntity.Task, error) {
	return *task, nil
}

func (l *taskRepo) GetAllTasks(filter *sharedModel.CriteriaFilter) ([]taskDto.TaskRecord, int64, error) {
	tasks := []taskDto.TaskRecord{}
	return tasks, 2, nil
}

func (l *taskRepo) GetTaskRecordById(taskId uint) (taskDto.TaskRecord, error) {
	task := taskDto.TaskRecord{}

	if taskId == 0 {
		return task, errors.New("not allowed param")
	}

	return task, nil
}

func (l *taskRepo) GetTaskById(taskId uint) (taskManagementEntity.Task, error) {
	task := taskManagementEntity.Task{}

	if taskId == 0 {
		return task, errors.New("not allowed param")
	}

	return task, nil
}

func (l *taskRepo) ChangeTaskStatus(taskId uint, status uint) error {

	if taskId == 0 {
		return errors.New("not allowed param")
	}

	return nil
}

func (l *taskRepo) GetAllTasksStatus() ([]taskManagementEntity.TaskStatus, error) {
	taskStatus := []taskManagementEntity.TaskStatus{}
	return taskStatus, nil
}

func (l *taskRepo) DeleteTask(taskId uint) error {
	if taskId == 0 {
		return errors.New("not allowed param")
	}

	return nil
}

func NewTaskRepo() taskManagementContract.TaskRepository {
	return &taskRepo{}
}
