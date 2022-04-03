package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"wefunder/service"
)

type DeckController struct {
	fileManager service.IFileManager
}

func NewDeckController(fm service.IFileManager) *DeckController {
	return &DeckController{
		fileManager: fm,
	}
}

func (r *DeckController) UploadDeck(c echo.Context) error {
	companyName := c.FormValue("companyName")
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fmt.Println(companyName)
	err = r.fileManager.PutFile(file)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "ok")
}

func (r *DeckController) GetDecks(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
