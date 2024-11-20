package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jonatasemanuel/echo-htmx/internal/models"
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
	res := map[string]interface{}{"name": all}
	return c.JSON(http.StatusOK, res)
}
