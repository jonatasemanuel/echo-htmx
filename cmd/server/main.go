package main

import (
	"time"

	"github.com/a-h/templ"
	"github.com/jonatasemanuel/templ-echo/internal/views"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", homeHandler)
	e.GET("/time", timeHandler)
	e.Logger.Fatal(e.Start(":3000"))
}

func homeHandler(c echo.Context) error {
	return render(c, views.Home("joey"))
}
func timeHandler(c echo.Context) error {
	return render(c, views.TimeComponent(time.Now()))
}

func notFoundHandler(c echo.Context) error {
	return render(c, views.NotFoundComponent())
}
func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}
