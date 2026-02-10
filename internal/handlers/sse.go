package handlers

import (
	"bytes"

	"github.com/labstack/echo/v4"
	"github.com/starfederation/datastar-go/datastar"

	appmw "webapp/internal/middleware"
	"webapp/internal/utils"
)

func SSE(c echo.Context) error {
	r := c.Request()
	w := c.Response().Writer
	log := appmw.Logger(c)

	cookie, err := c.Cookie("client_id")
	if err != nil {
		log.Warn("sse: no client_id cookie")
		return c.String(400, "No client_id cookie")
	}
	clientID := cookie.Value

	log.Debug("sse: client connected")

	sse := datastar.NewSSE(w, r)
	client := utils.Store.AddClient(clientID, sse)

	// Send initial body
	var buf bytes.Buffer
	data := utils.Store.GetTodos()

	if err := c.Echo().Renderer.Render(&buf, "body", data, c); err != nil {
		log.Error("sse: template error", "err", err)
		return c.NoContent(500)
	}
	sse.PatchElements(buf.String())
	log.Debug("sse: initial body sent", "todo_count", len(data))

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

			var buf bytes.Buffer
			data := utils.Store.GetTodos()

			if err := c.Echo().Renderer.Render(&buf, "body", data, c); err != nil {
				log.Error("sse: template error", "err", err)
				continue
			}

			if err := sse.PatchElements(buf.String()); err != nil {
				log.Error("sse: patch error", "err", err)
				return nil
			}
			log.Debug("sse: update sent", "todo_count", len(data))
		}
	}
}
