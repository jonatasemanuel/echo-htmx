package main

import (
	"github.com/a-h/templ"
	views "github.com/jonatasemanuel/templ-echo/internal/views/public"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", homeHandler)
	e.Logger.Fatal(e.Start(":3000"))
}

func homeHandler(c echo.Context) error {
	return render(c, views.Home("joey"))
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}
