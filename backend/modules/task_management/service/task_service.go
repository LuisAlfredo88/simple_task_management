package service

import (
	"errors"
	taskManagementContract "stm/modules/task_management/domain/contract"
	taskDto "stm/modules/task_management/domain/dto"
	taskManagementEntity "stm/modules/task_management/domain/entity"
	sharedModel "stm/shared/model"
)

type task struct {
	taskRepo taskManagementContract.TaskRepository
}

func (l *task) Save(task *taskManagementEntity.Task, userId string) (taskDto.TaskRecord, error) {
	// Setting logged user as the creator of the task
	if task.Id == 0 {
		task.CreatedById = userId
	} else {
		// Prevent that another user update the task if it's not the owner
		if task.CreatedById != userId {
			return taskDto.TaskRecord{}, errors.New("you are not allowed to edit this task")
		}
	}

	// Validating task information
	if err := task.Validate(); err != nil {
		return taskDto.TaskRecord{}, err
	}

	taskRecord, err := l.taskRepo.Save(task)

	if err != nil {
		return taskDto.TaskRecord{}, err
	}

	return l.taskRepo.GetTaskRecordById(taskRecord.Id)
}

func (l *task) GetAllTasks(filter *sharedModel.CriteriaFilter) ([]taskDto.TaskRecord, int64, error) {
	return l.taskRepo.GetAllTasks(filter)
}

func (l *task) GetTaskById(taskId uint) (taskManagementEntity.Task, error) {
	if taskId == 0 {
		return taskManagementEntity.Task{}, errors.New("not allowed param")
	}

	task, err := l.taskRepo.GetTaskById(taskId)

	if err != nil {
		return taskManagementEntity.Task{}, err
	}

	return task, err
}

func (l *task) ChangeTaskStatus(taskId uint, status uint) error {
	if taskId == 0 {
		return errors.New("not allowed param")
	}

	return l.taskRepo.ChangeTaskStatus(taskId, status)
}

func (l *task) GetAllTasksStatus() ([]taskManagementEntity.TaskStatus, error) {
	return l.taskRepo.GetAllTasksStatus()
}

func (l *task) DeleteTask(taskId uint) error {
	if taskId == 0 {
		return errors.New("not allowed param")
	}

	return l.taskRepo.DeleteTask(taskId)
}

func NewTaskService(
	taskRepo taskManagementContract.TaskRepository,
) taskManagementContract.TaskService {
	return &task{
		taskRepo: taskRepo,
	}
}
