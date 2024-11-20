package main

import (
	"log"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jonatasemanuel/echo-htmx/internal/database"
	"github.com/jonatasemanuel/echo-htmx/internal/handlers"
	"github.com/jonatasemanuel/echo-htmx/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// run() -> server start config
	dbConn, err := database.ConnectDB("./anime.db")
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer dbConn.DB.Close()
	e := echo.New()
	handlers.SessionManager = scs.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	handlers.SessionManager.Lifetime = 24 * time.Hour
	e.Use(echo.WrapMiddleware(handlers.SessionManager.LoadAndSave))

	// routes() -> call routes

	// --animes
	e.POST("/anime", handlers.PostAnime)
	e.GET("/anime", handlers.GetAllAnimes)

	// --guess
	e.GET("/final-score", handlers.FinalScore)
	e.GET("/", handlers.GetHome)
	e.POST("/", handlers.PostHomeHandler)

	// --counter
	e.GET("/count", handlers.GetHandler)
	e.POST("/count", handlers.PostHandler)

	// configs()
	models.New(dbConn.DB)

	e.Logger.Fatal(e.Start(":8080"))
}
