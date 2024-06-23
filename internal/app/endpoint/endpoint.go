package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-echo-rest-api/internal/app/model"
	"net/http"
)

type Service interface {
	DaysLeft() int64
	SaveCustomer(firstName, lastName, email, country, password string) error
	GetAllCustomers() ([]*model.Customer, error)
	DeleteCustomerByEmail(email string) error
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {
	d := e.s.DaysLeft()

	s := fmt.Sprintf("Days left: %d", d)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}

func (e *Endpoint) CreateCustomer(c echo.Context) error {
	customer := new(model.Customer)
	if err := c.Bind(customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	err := e.s.SaveCustomer(customer.FirstName, customer.LastName, customer.Email, customer.Country, customer.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save customer"})
	}

	return c.JSON(http.StatusCreated, customer)
}

func (e *Endpoint) GetAllCustomers(c echo.Context) error {
	customers, err := e.s.GetAllCustomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve customers"})
	}
	return c.JSON(http.StatusOK, customers)
}

func (e *Endpoint) DeleteCustomer(c echo.Context) error {
	type RequestBody struct {
		Email string `json:"email"`
	}

	reqBody := new(RequestBody)
	if err := c.Bind(reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if reqBody.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email is required"})
	}

	err := e.s.DeleteCustomerByEmail(reqBody.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete customer"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Customer deleted successfully"})
}
