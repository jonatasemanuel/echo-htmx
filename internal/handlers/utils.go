package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Contains(slice []map[string]interface{}, animeID int) bool {
	for _, item := range slice {
		if id, ok := item["ID"].(int); ok && id == animeID {
			return true
		}
	}
	return false
}

func Render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}
