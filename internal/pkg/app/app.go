package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-echo-rest-api/internal/app/endpoint"
	"go-echo-rest-api/internal/app/middleware"
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

	a.echo.Use(middleware.Verification)

	a.echo.GET("/status", a.e.Status)
	a.echo.POST("/customer", a.e.CreateCustomer)
	a.echo.GET("/customers", a.e.GetAllCustomers)

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
