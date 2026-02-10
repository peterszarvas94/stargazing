package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
)

func Add(c echo.Context) error {
	cookie, err := c.Cookie("client_id")
	if err != nil {
		log.Println("No client_id cookie")
	} else {
		log.Printf("Client %s added: %s", cookie.Value, c.FormValue("text"))
	}
	return c.NoContent(200)
}
