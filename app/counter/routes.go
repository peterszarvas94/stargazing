package counter

import (
	"html/template"

	"github.com/labstack/echo/v4"

	"webapp/internal/sse"
)

// New creates a new Counter controller and registers routes.
func New(e *echo.Echo) *Counter {
	ctr := &Counter{}

	// Parse feature template (contains both "index" and "app" blocks)
	ctr.tmpl = template.Must(template.ParseFiles(
		"app/counter/templates/index.html",
	))

	e.GET("/counter", ctr.Index)
	e.PATCH("/counter", ctr.Update)
	e.GET("/counter/sse", sse.Handler(ctr.tmpl, "app", ctr.Data))

	return ctr
}
