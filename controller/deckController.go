package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"wefunder/repository"
	"wefunder/service"
)

type DeckController struct {
	fileManager service.IFileManager
	repo *repository.DeckRepository
}

func NewDeckController(fm service.IFileManager, repo *repository.DeckRepository) *DeckController {
	return &DeckController{
		fileManager: fm,
		repo: repo,
	}
}

func (r *DeckController) UploadDeck(c echo.Context) error {
	companyName := c.FormValue("companyName")
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	files, err := r.fileManager.PutFile(file)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, files)
}

func (r *DeckController) GetDecks(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
