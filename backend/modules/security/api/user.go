package api

import (
	"fmt"
	"net/http"
	"strconv"

	securityEntity "stm/modules/security/domain/entity"

	sharedModel "stm/shared/model"

	"github.com/labstack/echo/v4"
)

func (l *SecurityApi) registerUserHandlers(api *echo.Group) {
	// Setting up security handlers
	api.POST("/users", l.register)
	api.GET("/users", l.getAllUsers)
	api.GET("/users/:id", l.getUserById)
	api.GET("/users/:name/exists", l.userExists)
	api.PATCH("/users/:user_id/:status", l.toogleUserStatus)
}

func (p *SecurityApi) getAllUsers(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	skipStr := c.QueryParam("skip")
	search := c.QueryParam("search")
	isActive := c.QueryParam("status")

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
			"isActive": isActive,
			"search":   search,
		},
	}

	users, totalRecords, _ := p.userService.GetAllUsers(&criteria)
	c.Response().Header().Set("X-TotalRecords", fmt.Sprintf("%d", totalRecords))
	return c.JSON(http.StatusOK, users)
}

func (p *SecurityApi) register(c echo.Context) error {
	userData := new(securityEntity.User)

	c.Bind(userData)

	user, err := p.userService.Save(userData)

	if err != nil {
		return c.JSON(http.StatusBadRequest, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

func (p *SecurityApi) getUserById(c echo.Context) error {
	userId := c.Param("id")

	users, err := p.userService.GetUserById(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, users)
}

func (p *SecurityApi) toogleUserStatus(c echo.Context) error {
	userId := c.Param("user_id")
	statusParam := c.Param("status")

	status, err := strconv.Atoi(statusParam)
	if err != nil {
		return c.JSON(http.StatusNoContent, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	err = p.userService.ToggleUserStatus(userId, int8(status))

	if err != nil {
		return c.JSON(http.StatusBadRequest, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, nil)
}

func (p *SecurityApi) userExists(c echo.Context) error {
	name := c.Param("name")

	type response struct {
		Exists bool `json:"exists"`
	}

	userExists, err := p.userService.UserExists(name)

	if err != nil {
		return c.JSON(http.StatusBadRequest, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Exists: userExists,
	})
}
