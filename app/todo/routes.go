package todo

import "github.com/labstack/echo/v4"

// New creates a new Todo controller and registers routes.
func New(e *echo.Echo) *Todo {
	t := &Todo{
		items: []Item{},
	}

	e.GET("/todo", t.Index)
	e.POST("/todo", t.Create)
	e.GET("/sse", t.SSE)

	return t
}
