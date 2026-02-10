package utils

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"webapp/internal/store"
)

var Store *store.Store

// EnsureClientID returns the client ID from the cookie, or generates a new one and sets the cookie.
func EnsureClientID(c echo.Context) string {
	cookie, err := c.Cookie("client_id")
	if err == nil {
		return cookie.Value
	}

	clientID := fmt.Sprintf("client_%d", time.Now().UnixNano())
	c.SetCookie(&http.Cookie{
		Name:     "client_id",
		Value:    clientID,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400,
	})
	return clientID
}

// GetClientID returns the client ID from the cookie or an error if not set.
func GetClientID(c echo.Context) (string, error) {
	cookie, err := c.Cookie("client_id")
	if err != nil {
		return "", errors.New("no client_id cookie")
	}
	return cookie.Value, nil
}

// RenderToString renders a template to a string.
func RenderToString(c echo.Context, name string, data any) (string, error) {
	var buf bytes.Buffer
	if err := c.Echo().Renderer.Render(&buf, name, data, c); err != nil {
		return "", err
	}
	return buf.String(), nil
}
