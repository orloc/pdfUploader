package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"wefunder/controller"
	"wefunder/service"
)

const filePath = "uploads"

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	  Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
	  Level: 5,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	  AllowOrigins: []string{"http://localhost:3000"},
	  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))


	fm := service.NewLocalFileManger(filePath)
	deckCtrl := controller.NewDeckController(fm)

	e.GET("/decks", deckCtrl.GetDecks)
	e.POST("/deck", deckCtrl.UploadDeck)

	e.Logger.Fatal(e.Start(":1323"))
}
