package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/alexedwards/scs/v2"
	"github.com/jonatasemanuel/echo-htmx/internal/database"
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
	end   int = 4
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
	animeName := charData["anime"]
	animeList := FetchData().AnimeList[start:end]

	slice := []string{}
	slice = append(slice, animeList...)
	if !contains(slice, animeName) {
		slice[0] = animeName
	}
	sort.Strings(slice)

	component := views.Home(strconv.Itoa(total.Count), charData, slice, done)
	return render(c, component)
}

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

	fmt.Println("")
	fmt.Println(end)
	fmt.Println("")
	animesQtd := len(FetchData().AnimeList)
	if end < animesQtd {
		start += 4
		end += 4
		char++
		done--
	}
	fmt.Println("")
	fmt.Println(end)
	fmt.Println("")
	if end > animesQtd {
		start = 0
		end = 4
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
	if err := database.Run(); err != nil {
		log.Fatal(err)
	}
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
