package health

import "github.com/labstack/echo/v4"

// New creates a new Health controller and registers routes.
func New(e *echo.Echo) *Health {
	h := &Health{}

	e.GET("/health", h.Check)

	return h
}
