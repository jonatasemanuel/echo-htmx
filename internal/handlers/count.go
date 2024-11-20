package handlers

import (
	"github.com/alexedwards/scs/v2"
	views "github.com/jonatasemanuel/echo-htmx/internal/views/public"
	"github.com/labstack/echo/v4"
)

type GlobalState struct {
	Count int
}

var global GlobalState
var SessionManager *scs.SessionManager

func GetHandler(c echo.Context) error {
	userCount := SessionManager.GetInt(c.Request().Context(), "count")
	component := views.Page(global.Count, userCount)
	return Render(c, component)
}

func PostHandler(c echo.Context) error {
	if c.FormValue("global") != "" {
		global.Count++
	}
	if c.FormValue("session") != "" {
		currentCount := SessionManager.GetInt(c.Request().Context(), "count")
		SessionManager.Put(c.Request().Context(), "count", currentCount+1)
	}
	return GetHandler(c)
}
