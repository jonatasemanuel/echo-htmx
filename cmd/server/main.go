package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jonatasemanuel/echo-htmx/internal/database"
	"github.com/jonatasemanuel/echo-htmx/internal/handlers"
	"github.com/jonatasemanuel/echo-htmx/internal/models"
	views "github.com/jonatasemanuel/echo-htmx/internal/views/public"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	return handlers.Render(c, component)
}

func getHome(c echo.Context) error {
	charData := FetchData().Char[char]
	animeName := charData["anime"]
	animeList := FetchData().AnimeList[start:end]

	slice := []string{}
	slice = append(slice, animeList...)
	if !handlers.Contains(slice, animeName) {
		slice[0] = animeName
	}
	sort.Strings(slice)

	component := views.Home(strconv.Itoa(total.Count), charData, slice, done)
	return handlers.Render(c, component)
}

func postHomeHandler(c echo.Context) error {
	if c.FormValue("total") != "" {
		total.Count++
	}

	animesQtd := len(FetchData().AnimeList) - 1
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

func main() {
	// run() -> server start config
	dbConn, err := database.ConnectDB("./anime.db")
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer dbConn.DB.Close()
	e := echo.New()
	handlers.SessionManager = scs.New()

	// routes() -> call routes
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	handlers.SessionManager.Lifetime = 24 * time.Hour
	e.Use(echo.WrapMiddleware(handlers.SessionManager.LoadAndSave))

	e.POST("/anime", func(c echo.Context) error {
		var animeData models.Anime

		err := json.NewDecoder(c.Request().Body).Decode(&animeData)
		if err != nil {
			log.Print("error to create")
		}

		animeCreated, err := anime.CreateAnime(animeData)
		if err != nil {
			log.Print("error to save")
		}

		return c.JSON(http.StatusOK, animeCreated)
	})
	e.GET("/anime", func(c echo.Context) error {
		all, err := anime.ListAnimes()
		if err != nil {
			log.Fatal(err)
		}
		res := map[string]interface{}{"name": all}
		return c.JSON(http.StatusOK, res)
	})

	e.GET("/final-score", finalScore)
	e.GET("/", getHome)
	e.POST("/", postHomeHandler)

	// --counter
	e.GET("/count", handlers.GetHandler)
	e.POST("/count", handlers.PostHandler)

	// configs()
	models.New(dbConn.DB)

	e.Logger.Fatal(e.Start(":8080"))

}
