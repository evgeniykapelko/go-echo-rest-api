package middleware

import (
	"github.com/labstack/echo/v4"
)

func InitMiddlewares(e *echo.Echo) {
	e.Use(Verification)
}
