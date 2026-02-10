package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	// Set client_id cookie if not exists
	_, err := c.Cookie("client_id")
	if err != nil {
		cookie := &http.Cookie{
			Name:     "client_id",
			Value:    fmt.Sprintf("client_%d", time.Now().UnixNano()),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   86400,
		}
		c.SetCookie(cookie)
	}
	return c.Render(http.StatusOK, "index.html", nil)
}
