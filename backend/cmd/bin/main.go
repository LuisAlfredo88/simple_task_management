package main

import (
	"net/http"

	"stm/initialize/api"
	"stm/initialize/db/sqlite"
	"stm/initialize/repository"
	"stm/initialize/service"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {

	app := fx.New(
		fx.Provide(
			echo.New,
			http.NewServeMux,
			sqlite.NewSqlInstance,
			api.NewHTTPServer,
		),
		repository.Module,
		service.Module,
		api.Module,
		fx.Invoke(api.RegisterRoutes),
	)

	app.Run()
}
