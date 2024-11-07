package main

import (
	"time"

	"github.com/a-h/templ"
	"github.com/alexedwards/scs/v2"
	views "github.com/jonatasemanuel/echo-htmx/internal/views/public"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
func homeHandler(c echo.Context) error {
	return render(c, views.Home("joey"))
}
func timeHandler(c echo.Context) error {
	return render(c, views.TimeComponent(time.Now()))
} */

type GlobalState struct {
	Count int
}
type Quest struct {
	Char   string
	Animes [4]string
}

var global GlobalState
var sessionManager *scs.SessionManager

func getHandler(c echo.Context) error {
	userCount := sessionManager.GetInt(c.Request().Context(), "count")
	component := views.Page(global.Count, userCount)
	return render(c, component)
}

func postHandler(c echo.Context) error {
	if c.FormValue("global") != "" {
		global.Count++
	}
	if c.FormValue("session") != "" {
		currentCount := sessionManager.GetInt(c.Request().Context(), "count")
		sessionManager.Put(c.Request().Context(), "count", currentCount+1)
	}
	return getHandler(c)
}

func getHome(c echo.Context) error {
	quest := Quest{
		Char:   "Ace",
		Animes: [4]string{"One Piece", "Naruto", "Bleach", "X-men"},
	}

	return render(c, views.Home("guess the anime", quest.Char, quest.Animes))
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	e.Use(echo.WrapMiddleware(sessionManager.LoadAndSave))

	e.GET("/", getHome)
	e.GET("/count", getHandler)
	e.POST("/count", postHandler)
	e.Logger.Fatal(e.Start(":8080"))

}
