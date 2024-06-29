package api

import (
	taskManagementService "stm/modules/task_management/domain/contract"

	"github.com/labstack/echo/v4"
)

type TaskManagementApi struct {
	taskService taskManagementService.TaskService
	api         *echo.Echo
}

func NewTaskManagementApi(
	taskService taskManagementService.TaskService,
	api *echo.Echo,
) *TaskManagementApi {
	rest_api := &TaskManagementApi{
		taskService: taskService,
		api:         api,
	}

	return rest_api
}

func (l *TaskManagementApi) Register() {
	taskManagement := l.api.Group("/tasks_management")
	//Registering task handlers
	l.registerTaskHandlers(taskManagement)
}
