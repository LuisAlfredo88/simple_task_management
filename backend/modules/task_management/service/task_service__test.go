package service

import (
	"errors"
	"testing"

	taskDto "stm/modules/task_management/domain/dto"
	taskManagementEntity "stm/modules/task_management/domain/entity"
	MockTaskManagementRepo "stm/modules/task_management/repository/mock"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	mockTaskManagementRepo := MockTaskManagementRepo.NewTaskRepo()
	taskService := &task{
		taskRepo: mockTaskManagementRepo,
	}

	userId := "user123"
	task := taskManagementEntity.Task{
		Id:          1,
		Title:       "This is the title",
		Description: "This is the description",
		StatusId:    1,
		CreatedById: userId,
	}

	data := []map[string]interface{}{
		{
			"name":          "Title validation",
			"ExpectedError": "must specify task title",
			"Data": func() *taskManagementEntity.Task {
				t := task
				t.Title = ""
				return &t
			},
		},

		{
			"name":          "Description validation",
			"ExpectedError": "must specify task description",
			"Data": func() *taskManagementEntity.Task {
				t := task
				t.Description = ""
				return &t
			},
		},

		{
			"name":          "Status validation",
			"ExpectedError": "must specify status",
			"Data": func() *taskManagementEntity.Task {
				t := task
				t.StatusId = 0
				return &t
			},
		},

		{
			"name":          "Created By validation",
			"ExpectedError": "must specify created by",
			"Data": func() *taskManagementEntity.Task {
				t := task
				t.Id = 0
				t.CreatedById = ""
				userId = ""
				return &t
			},
		},

		{
			"name":          "Edition validation",
			"ExpectedError": "you are not allowed to edit this task",
			"Data": func() *taskManagementEntity.Task {
				t := task
				t.Id = 1
				t.CreatedById = "user123"
				userId = "user345"
				return &t
			},
		},
	}

	t.Run("validation error", func(t *testing.T) {

		for _, d := range data {
			name := d["name"].(string)
			expected := d["ExpectedError"].(string)
			info, _ := d["Data"].(func() *taskManagementEntity.Task)

			t.Run(name, func(t *testing.T) {
				validationError := errors.New(expected)
				_, err := taskService.Save(info(), userId)

				assert.Error(t, err)
				assert.Equal(t, validationError, err)
			})
		}
	})

	t.Run("save method must return a TaskRecord instance", func(t *testing.T) {
		result, err := taskService.Save(&task, "user123")
		assert.Equal(t, err, nil)
		assert.Equal(t, result, taskDto.TaskRecord{})
	})
}
