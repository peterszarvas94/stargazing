package home

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

type HomeData struct {
	Title string
}

type Home struct {
	tmpl *template.Template
}

// Data returns data for rendering.
func (h *Home) Data() any {
	return HomeData{
		Title: "Stargazing",
	}
}

// New creates a new Home controller and registers routes.
func New(e *echo.Echo) *Home {
	h := &Home{}

	// Parse home-specific template (full page, no shared templates)
	h.tmpl = template.Must(template.ParseFiles(
		"app/home/templates/index.html",
	))

	e.GET("/", h.Index)

	return h
}
