package service

import (
	taskManagementContract "stm/modules/task_management/domain/contract"
	taskDto "stm/modules/task_management/domain/dto"
	taskManagementEntity "stm/modules/task_management/domain/entity"
	sharedModel "stm/shared/model"
)

type task struct {
	taskRepo           taskManagementContract.TaskRepository
	taskManagementRepo taskManagementContract.TaskRepository
}

func (l *task) Save(task *taskManagementEntity.Task, userId string) (taskDto.TaskRecord, error) {
	// Setting logged user as the creator of the task
	if task.Id == 0 {
		task.CreatedById = userId
	}

	if err := task.Validate(); err != nil {
		return taskDto.TaskRecord{}, err
	}

	taskRecord, err := l.taskRepo.Save(task)

	if err != nil {
		return taskDto.TaskRecord{}, err
	}

	return l.taskManagementRepo.GetTaskRecordById(taskRecord.Id)
}

func (l *task) GetAllTasks(filter *sharedModel.CriteriaFilter) ([]taskDto.TaskRecord, int64, error) {
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

func (l *task) GetAllTasksStatus() ([]taskManagementEntity.TaskStatus, error) {
	return l.taskRepo.GetAllTasksStatus()
}

func (l *task) DeleteTask(taskId uint) error {
	return l.taskRepo.DeleteTask(taskId)
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
