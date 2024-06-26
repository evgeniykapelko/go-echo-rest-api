package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-echo-rest-api/internal/app/endpoint"
	"go-echo-rest-api/internal/app/middleware"
	"go-echo-rest-api/internal/app/routes"
	"go-echo-rest-api/internal/app/service"
	"log"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	middleware.InitMiddlewares(a.echo)

	routes.InitRoutes(a.echo, a.e)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server running")

	err := a.echo.Start(":8081")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
