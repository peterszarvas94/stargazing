package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Index renders the home page with links to examples.
func (h *Home) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
