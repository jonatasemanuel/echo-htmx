package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jonatasemanuel/echo-htmx/internal/models"
	"github.com/jonatasemanuel/echo-htmx/internal/views/public"
	"github.com/labstack/echo/v4"
)

func PostAnime(c echo.Context) error {
	var animeData models.Anime

	err := json.NewDecoder(c.Request().Body).Decode(&animeData)
	if err != nil {
		log.Fatal("error to create")
	}

	animeCreated, err := anime.CreateAnime(animeData)
	if err != nil {
		log.Fatal("error to save")
	}

	return c.JSON(http.StatusOK, animeCreated)

}
func GetAllAnimes(c echo.Context) error {
	all, err := anime.ListAnimes()
	if err != nil {
		log.Fatal(err)
	}
	animeList := make([]map[string]interface{}, 0)
	for _, anime := range all {
		animeList = append(animeList, map[string]interface{}{
			"ID":   anime.ID,
			"Name": anime.Name,
		})

	}
	component := views.Anime(animeList)
	return Render(c, component)
}
