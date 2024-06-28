package api

import (
	"net/http"

	securityDto "stm/modules/security/domain/dto"

	sharedModel "stm/shared/model"

	"github.com/labstack/echo/v4"
)

func (l *SecurityApi) registerLoginHandlers(api *echo.Group) {
	// Setting up security handlers
	api.POST("/authenticate", l.authenticate)
	api.POST("/token_refresh", l.refreshToken)
}

func (p *SecurityApi) authenticate(c echo.Context) error {
	userData := securityDto.LoginData{}
	c.Bind(&userData)

	logedUser, err := p.securityService.Authenticate(userData)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, sharedModel.ResponseMessage{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, logedUser)
}

func (p *SecurityApi) refreshToken(c echo.Context) error {
	userId, _ := c.Get("userId").(string)

	token, _ := p.securityService.RefreshToken(userId)

	responseData := map[string]string{
		"token": token,
	}

	return c.JSON(http.StatusOK, responseData)
}
