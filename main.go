package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"wefunder/controller"
	"wefunder/repository"
	"wefunder/service"
	"wefunder/util"
)

const filePath = "uploads"

func main() {
	env := util.LoadEnv()
	pg, pgCtx := util.AcquirePostgresPool(env.DBUrl)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
	  Level: 5,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	  AllowOrigins: []string{"http://localhost:3000"},
	  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))


	deckRepo := repository.NewDeckRepository(pgCtx, pg)

	fm := service.NewLocalFileManger(filePath)
	deckCtrl := controller.NewDeckController(fm, deckRepo)

	e.GET("/decks", deckCtrl.GetDecks)
	e.POST("/deck", deckCtrl.UploadDeck)

	e.Logger.Fatal(e.Start(":3030"))
}
