package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

func Verification(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")

		if strings.EqualFold(val, "secret") {
			log.Println("Using secret")
		}

		err := next(ctx)

		if err != nil {
			return err
		}

		return nil
	}
}
