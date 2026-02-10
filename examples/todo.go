package examples

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/starfederation/datastar-go/datastar"

	appmw "webapp/internal/middleware"
	"webapp/internal/utils"
)

func TodoPage(c echo.Context) error {
	clientID := utils.EnsureClientID(c)

	log := appmw.Logger(c)
	data := utils.Store.GetTodos()

	log.Debug("todo: rendering", "client_id", clientID, "todo_count", len(data))
	return c.Render(http.StatusOK, "todo", data)
}

type AddTodoSignals struct {
	Todo string `json:"todo"`
}

func AddTodo(c echo.Context) error {
	log := appmw.Logger(c)

	var signals AddTodoSignals
	if err := datastar.ReadSignals(c.Request(), &signals); err != nil {
		log.Warn("todo: failed to read signals", "err", err)
		return c.NoContent(400)
	}

	if signals.Todo == "" {
		log.Debug("todo: empty text")
		return c.NoContent(200)
	}

	todoCount := utils.Store.AddTodo(signals.Todo)
	log.Debug("todo: created", "text", signals.Todo, "todo_count", todoCount)

	cookie, err := c.Cookie("client_id")
	if err != nil {
		log.Warn("todo: no client_id cookie")
		return c.NoContent(200)
	}

	if err := utils.Store.SignalClient(cookie.Value); err != nil {
		log.Warn("todo: signal failed", "err", err)
	} else {
		log.Debug("todo: signal sent")
	}

	return c.NoContent(200)
}
