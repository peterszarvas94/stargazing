package home

import "github.com/labstack/echo/v4"

// New creates a new Home controller and registers routes.
func New(e *echo.Echo) *Home {
	h := &Home{}

	e.GET("/", h.Index)

	return h
}
