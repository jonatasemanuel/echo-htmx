package main

import (
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/alexedwards/scs/v2"
	views "github.com/jonatasemanuel/echo-htmx/internal/views/public"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ScoreState struct {
	Count int
}
type GlobalState struct {
	Count int
}
type Quest struct {
	Chars  []string
	Animes []string
}

var (
	char  int = 0
	start int = 0
	end   int = 4
)
var global GlobalState
var total ScoreState
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
	// get from some repo(api) or create a api(pkg) fetching by another
	slice := FetchQuestData().Animes[start:end]
	component := views.Home(strconv.Itoa(total.Count), FetchQuestData().Chars[char], slice)
	return render(c, component)
}

func postHomeHandler(c echo.Context) error {
	if c.FormValue("total") != "" {
		total.Count++
	}
	start += 4
	end += 4
	char++
	return getHome(c)
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func main() {
	e := echo.New()
	sessionManager = scs.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	sessionManager.Lifetime = 24 * time.Hour
	e.Use(echo.WrapMiddleware(sessionManager.LoadAndSave))

	e.GET("/", getHome)
	e.POST("/", postHomeHandler)

	// --counter
	e.GET("/count", getHandler)
	e.POST("/count", postHandler)
	e.Logger.Fatal(e.Start(":8080"))

}
