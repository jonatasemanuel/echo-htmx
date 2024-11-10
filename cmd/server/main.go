package main

import (
	"net/http"
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
	end   int = 3
	done  int = 15
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

func finalScore(c echo.Context) error {
	component := views.FinalScore(strconv.Itoa(total.Count))
	return render(c, component)
}

func getHome(c echo.Context) error {
	// get from some repo(api) or create a api(pkg) fetching by another
	// need turn into a map
	slice := FetchQuestData().Animes[start : end+1]
	component := views.Home(strconv.Itoa(total.Count), FetchQuestData().Chars[char], slice, done)
	return render(c, component)
}

func postHomeHandler(c echo.Context) error {
	if c.FormValue("total") != "" {
		total.Count++
	}

	if end < len(FetchQuestData().Animes) {
		start += 4
		end += 4
		char++
		done--
	}
	if end > len(FetchQuestData().Animes) {
		start = 0
		end = 4
		char = 0
		c.Response().Header().Set("HX-Redirect", "/final-score")
		return c.NoContent(http.StatusNoContent)
	}
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

	e.GET("/final-score", finalScore)
	e.GET("/", getHome)
	e.POST("/", postHomeHandler)

	// --counter
	e.GET("/count", getHandler)
	e.POST("/count", postHandler)
	e.Logger.Fatal(e.Start(":8080"))

}
