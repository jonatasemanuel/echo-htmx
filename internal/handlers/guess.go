package handlers

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/jonatasemanuel/echo-htmx/cmd"
	"github.com/jonatasemanuel/echo-htmx/internal/models"
	views "github.com/jonatasemanuel/echo-htmx/internal/views/public"
	"github.com/labstack/echo/v4"
)

type ScoreState struct {
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

var total ScoreState
var anime models.Anime

func finalScore(c echo.Context) error {
	component := views.FinalScore(strconv.Itoa(total.Count))
	return Render(c, component)
}

func getHome(c echo.Context) error {
	charData := cmd.FetchData().Char[char]
	animeName := charData["anime"]
	animeList := cmd.FetchData().AnimeList[start:end]

	slice := []string{}
	slice = append(slice, animeList...)
	if !Contains(slice, animeName) {
		slice[0] = animeName
	}
	sort.Strings(slice)

	component := views.Home(strconv.Itoa(total.Count), charData, slice, done)
	return Render(c, component)
}

func postHomeHandler(c echo.Context) error {
	if c.FormValue("total") != "" {
		total.Count++
	}

	animesQtd := len(cmd.FetchData().AnimeList) - 1
	if end < animesQtd {
		start += 4
		end += 4
		char++
		done--
	}
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
