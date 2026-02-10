package todo

import (
	"html/template"

	"github.com/labstack/echo/v4"

	"webapp/internal/sse"
)

// New creates a new Todo controller and registers routes.
func New(e *echo.Echo) *Todo {
	t := &Todo{
		items: []Item{},
	}

	// Parse feature template (contains both "index" and "app" blocks)
	t.tmpl = template.Must(template.ParseFiles(
		"app/todo/templates/index.html",
	))

	e.GET("/todo", t.Index)
	e.POST("/todo", t.Create)
	e.GET("/todo/sse", sse.Handler(t.tmpl, "app", t.Data))

	return t
}
