package main

import (
	"net/http"
	"sort"
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

type Data struct {
	Char      []map[string]string
	AnimeList []string
}

var (
	slice []string
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
	charData := FetchData().Char[char]

	slice = []string{charData["anime"]}
	animeList := FetchData().AnimeList
	for _, anime := range animeList {
		if len(slice) >= 4 {
			break
		}
		if anime != charData["anime"] && !contains(slice, anime) {
			slice = append(slice, anime)
		}

	}

	sort.Strings(slice)

	component := views.Home(strconv.Itoa(total.Count), charData, slice, done)
	return render(c, component)
}

// Helper function to check if a slice contains a specific item.
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func postHomeHandler(c echo.Context) error {
	if c.FormValue("total") != "" {
		total.Count++
	}

	if done > 0 {
		char++
		done--
	}
	if done == 0 {
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
