package routes

import (
	"github.com/labstack/echo/v4"
	"go-echo-rest-api/internal/app/endpoint"
)

func InitRoutes(e *echo.Echo, endpoint *endpoint.Endpoint) {
	e.GET("/status", endpoint.Status)
	e.POST("/create", endpoint.CreateCustomer)
	e.GET("/customers", endpoint.GetAllCustomers)
	e.DELETE("/delete", endpoint.DeleteCustomer)
	e.GET("/customer", endpoint.GetCustomerByEmail)
	// e.POST("/update", endpoint.UpdateCustomerById) // Update customer by ID
}
