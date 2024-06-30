package api

import (
	"net/http"
	"stm/shared/utility/jwt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type API struct {
	db *gorm.DB
}

func NewHTTPServer(lc fx.Lifecycle, srv *http.ServeMux, db *gorm.DB) *API {
	return &API{
		db: db,
	}
}

func UserHeaderMiddleware(next echo.HandlerFunc, db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() == "/security/authenticate" {
			return next(c)
		}

		authData := c.Request().Header.Get("Authorization")

		if authData == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid request")
		}

		token := strings.Fields(authData)[1]
		jwtData, err := jwt.Parse(token)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid request")
		}

		c.Set("userId", jwtData["uid"])

		return next(c)
	}
}

func (s *API) Start(e *echo.Echo) {
	e.Use(middleware.CORS())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"X-TotalRecords"},
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return UserHeaderMiddleware(next, s.db)
	})

	e.Start(":5000")
}
