package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	appmw "webapp/internal/middleware"
	"webapp/internal/utils"
)

func Index(c echo.Context) error {
	var clientID string
	isNew := false

	cookie, err := c.Cookie("client_id")
	if err != nil {
		clientID = fmt.Sprintf("client_%d", time.Now().UnixNano())
		c.SetCookie(&http.Cookie{
			Name:     "client_id",
			Value:    clientID,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   86400,
		})
		isNew = true
	} else {
		clientID = cookie.Value
	}

	log := appmw.Logger(c)
	data := utils.Store.GetTodos()

	log.Debug("index: rendering", "todo_count", len(data), "new_client", isNew)
	return c.Render(http.StatusOK, "index", data)
}
