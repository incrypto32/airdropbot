package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// Not found handler
	echo.NotFoundHandler = func(c echo.Context) error {
		// Render your 404 page
		return c.String(http.StatusNotFound, "404 Not found")
	}

	// Middlewares
	e.Use(middleware.Recover())

	return e
}
