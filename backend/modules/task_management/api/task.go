package api

import (
	"fmt"
	"net/http"
	"strconv"

	taskManagementEntity "stm/modules/task_management/domain/entity"

	sharedModel "stm/shared/model"

	"github.com/labstack/echo/v4"
)

func (l *TaskManagementApi) registerTaskHandlers(api *echo.Group) {
	// Setting up task handlers
	api.POST("/tasks", l.register)
	api.GET("/tasks", l.getAllTasks)
	api.GET("/task_status", l.GetAllTasksStatus)
	api.GET("/tasks/:id", l.getTaskById)
	api.PATCH("/tasks/:task_id/:status", l.changeTaskStatus)
}

func (p *TaskManagementApi) getAllTasks(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	skipStr := c.QueryParam("skip")
	status := c.QueryParam("status")
	search := c.QueryParam("search")
	createdBy := c.QueryParam("createdBy")
	assignedTo := c.QueryParam("assignedTo")

	limit, err := strconv.Atoi(limitStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	skip, err := strconv.Atoi(skipStr)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	criteria := sharedModel.CriteriaFilter{
		Limit: int32(limit),
		Skip:  int32(skip),
		Filters: map[string]interface{}{
			"status":     status,
			"search":     search,
			"createdBy":  createdBy,
			"assignedTo": assignedTo,
		},
	}

	if err := c.Bind(&criteria); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tasks, totalRecords, _ := p.taskService.GetAllTasks(&criteria)
	c.Response().Header().Set("X-TotalRecords", fmt.Sprintf("%d", totalRecords))
	return c.JSON(http.StatusOK, tasks)
}

func (p *TaskManagementApi) register(c echo.Context) error {
	userId, _ := c.Get("userId").(string)
	taskData := new(taskManagementEntity.Task)

	c.Bind(taskData)

	task, err := p.taskService.Save(taskData, userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, task)
}

func (p *TaskManagementApi) getTaskById(c echo.Context) error {
	taskIdParam := c.Param("id")
	taskId, err := strconv.Atoi(taskIdParam)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tasks, err := p.taskService.GetTaskById(uint(taskId))

	if err != nil {
		return c.JSON(http.StatusBadRequest, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (p *TaskManagementApi) changeTaskStatus(c echo.Context) error {
	taskIdParam := c.Param("task_id")
	statusParam := c.Param("status")

	status, err := strconv.Atoi(statusParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	taskId, err := strconv.Atoi(taskIdParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	err = p.taskService.ChangeTaskStatus(uint(taskId), uint(status))

	if err != nil {
		return c.JSON(http.StatusBadRequest, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, nil)
}

func (p *TaskManagementApi) GetAllTasksStatus(c echo.Context) error {
	taskStatus, err := p.taskService.GetAllTasksStatus()

	if err != nil {
		return c.JSON(http.StatusBadRequest, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, taskStatus)
}
