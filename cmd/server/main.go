package main

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"webapp/internal/handlers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	e.Renderer = &TemplateRenderer{templates: tmpl}

	e.GET("/", handlers.Index)
	e.POST("/add", handlers.Add)

	log.Println("Server starting on :8080")
	e.Start(":8080")
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
