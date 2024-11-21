package handlers

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/jonatasemanuel/echo-htmx/internal/models"
	views "github.com/jonatasemanuel/echo-htmx/internal/views/public"
	"github.com/labstack/echo/v4"
)

type ScoreState struct {
	Count int
}

var (
	char  int = 0
	start int = 0
	end   int = 4
	done  int = 15
)

var total ScoreState

var anime models.Anime
var character models.Character

func FinalScore(c echo.Context) error {
	component := views.FinalScore(strconv.Itoa(total.Count))
	return Render(c, component)
}

func GetHome(c echo.Context) error {
	charData, err := character.GetCharByID(char)
	for err != nil {
		char++
		charData, err = character.GetCharByID(char)
	}

	allAnimes, err := anime.ListAnimes()
	if err != nil {
		log.Fatal(err)
	}

	var animeName string
	animeList := make([]map[string]interface{}, 0)
	for _, anime := range allAnimes {
		animeList = append(animeList, map[string]interface{}{
			"ID":   anime.ID,
			"Name": anime.Name,
		})
		if anime.ID == charData.Anime {
			animeName = anime.Name
		}

	}

	slice := make([]map[string]interface{}, 0)
	for _, anime := range animeList[start:end] {
		slice = append(slice, anime)
	}

	if !Contains(slice, charData.Anime) {
		// rand number 1?-4 to index result
		slice[rand.Intn(4)] = map[string]interface{}{
			"ID":   charData.Anime,
			"Name": animeName,
		}
	}

	component := views.Home(strconv.Itoa(total.Count), charData, slice, done)
	return Render(c, component)
}

func PostHomeHandler(c echo.Context) error {
	if c.FormValue("total") != "" {
		total.Count++
	}

	allAnimes, err := anime.ListAnimes()
	if err != nil {
		log.Fatal(err)
	}

	animesQtd := len(allAnimes) - 1
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
	return GetHome(c)
}
