package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/starfederation/datastar-go/datastar"

	appmw "webapp/internal/middleware"
	"webapp/internal/utils"
)

type createSignals struct {
	Todo string `json:"todo"`
}

// Index renders the todo list page.
func (t *Todo) Index(c echo.Context) error {
	clientID := utils.EnsureClientID(c)

	log := appmw.Logger(c)
	data := t.All()

	log.Debug("todo.index", "client_id", clientID, "todo_count", len(data))
	return c.Render(http.StatusOK, "todo", data)
}

// Create handles adding a new todo item.
func (t *Todo) Create(c echo.Context) error {
	log := appmw.Logger(c)

	var signals createSignals
	if err := datastar.ReadSignals(c.Request(), &signals); err != nil {
		log.Warn("todo.create: failed to read signals", "err", err)
		return c.NoContent(400)
	}

	if signals.Todo == "" {
		log.Debug("todo.create: empty text")
		return c.NoContent(200)
	}

	todoCount := t.Add(signals.Todo)
	log.Debug("todo.create", "text", signals.Todo, "todo_count", todoCount)

	clientID, err := utils.GetClientID(c)
	if err != nil {
		log.Warn("todo.create: no client_id cookie")
		return c.NoContent(200)
	}

	if err := utils.Store.SignalClient(clientID); err != nil {
		log.Warn("todo.create: signal failed", "err", err)
	} else {
		log.Debug("todo.create: signal sent")
	}

	return c.NoContent(200)
}

// SSE handles the Server-Sent Events connection for real-time updates.
func (t *Todo) SSE(c echo.Context) error {
	r := c.Request()
	w := c.Response().Writer
	log := appmw.Logger(c)

	clientID, err := utils.GetClientID(c)
	if err != nil {
		log.Warn("sse: no client_id cookie")
		return c.String(400, "No client_id cookie")
	}

	log.Debug("sse: client connected")

	sse := datastar.NewSSE(w, r)
	client := utils.Store.AddClient(clientID, sse)

	// Send initial body
	data := t.All()
	html, err := utils.RenderToString(c, "app", data)
	if err != nil {
		log.Error("sse: template error", "err", err)
		return c.NoContent(500)
	}
	if err := sse.PatchElements(html); err != nil {
		log.Error("sse: initial patch error", "err", err)
		return c.NoContent(500)
	}
	log.Debug("sse: initial app sent", "todo_count", len(data))

	// Cleanup on disconnect
	defer func() {
		utils.Store.RemoveClient(clientID)
		log.Debug("sse: client disconnected")
	}()

	// Wait for signals and send updates
	for {
		select {
		case <-r.Context().Done():
			log.Debug("sse: context done")
			return nil
		case _, ok := <-client.Signals:
			if !ok {
				log.Debug("sse: signal channel closed")
				return nil
			}

			data := t.All()
			html, err := utils.RenderToString(c, "app", data)
			if err != nil {
				log.Error("sse: template error", "err", err)
				continue
			}

			if err := sse.PatchElements(html); err != nil {
				log.Error("sse: patch error", "err", err)
				return nil
			}
			log.Debug("sse: update sent", "todo_count", len(data))
		}
	}
}
