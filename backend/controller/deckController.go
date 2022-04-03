package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
	"strings"
	"wefunder/backend/entity"
	"wefunder/backend/repository"
	"wefunder/backend/service"
)

type response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Error string `json:"error"`
}

type DeckController struct {
	fileManager service.IFileManager
	repo        *repository.DeckRepository
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

	cName := strings.TrimSpace(companyName)
	exists, err := r.repo.Exists(cName)
	if err != nil {
		return err
	}

	if exists != nil {
		return c.JSON(http.StatusConflict, response{
			Status: http.StatusConflict,
			Error:  "Company already exists",
		})
	}


	files, err := r.fileManager.PutFile(file)
	if err != nil {
		return err
	}

	uuid := uuid.New()
	deck := new(entity.Deck)
	deck.CompanyName = cName
	deck.Uuid = uuid.String()
	deck.Images = files

	err = r.repo.CreateDeck(deck)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, response{
		Status: http.StatusCreated,
		Data:   deck,
		Error:  "",
	})
}

func (r *DeckController) GetDecks(c echo.Context) error {
	decks, err := r.repo.LoadDecks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: http.StatusOK,
		Data:   decks,
	})
}
