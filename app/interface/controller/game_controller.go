package controller

import (
	"net/http"
	"strconv"

	"github.com/egorkos/minesweeper/app/domain/model"
	"github.com/egorkos/minesweeper/app/registry"
	"github.com/egorkos/minesweeper/app/usecase"
	"github.com/gin-gonic/gin"
)

const (
	IdMustBeNumeric = "The ID must be numeric"
)

type square struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

func CreateGame(c *gin.Context) {
	var newGame model.Game
	err := c.BindJSON(&newGame)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = newGame.Validate()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctn := c.MustGet("ctn").(*registry.Container)
	useCase := ctn.Resolve("game-usecase").(usecase.GameUsecase)
	newGame, apiError := useCase.StartGame(newGame)

	if apiError != nil {
		c.AbortWithStatusJSON(apiError.Status, apiError.Error())
		return
	}

	c.JSON(http.StatusCreated, newGame)
	return
}

func GetGame(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, IdMustBeNumeric)
	}

	ctn := c.MustGet("ctn").(*registry.Container)
	useCase := ctn.Resolve("game-usecase").(usecase.GameUsecase)

	game, apiError := useCase.FindByID(ID)
	if apiError != nil {
		c.String(apiError.Status, apiError.Error())
		return
	}

	c.JSON(http.StatusOK, game)
	return
}

func ListGames(c *gin.Context) {
	ctn := c.MustGet("ctn").(*registry.Container)
	useCase := ctn.Resolve("game-usecase").(usecase.GameUsecase)

	games, apiError := useCase.FindAll()
	if apiError != nil {
		c.AbortWithStatusJSON(apiError.Status, apiError.Error())
		return
	}

	c.JSON(http.StatusOK, games)
	return
}

func Reveal(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, IdMustBeNumeric)
	}

	var square square
	err = c.BindJSON(&square)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctn := c.MustGet("ctn").(*registry.Container)
	useCase := ctn.Resolve("game-usecase").(usecase.GameUsecase)

	game, apiError := useCase.Reveal(ID, square.Row, square.Col)
	if apiError != nil {
		c.AbortWithStatusJSON(apiError.Status, apiError.Error())
		return
	}

	c.JSON(http.StatusOK, game)
	return
}

func Flag(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, IdMustBeNumeric)
	}

	var square square
	err = c.BindJSON(&square)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctn := c.MustGet("ctn").(*registry.Container)
	useCase := ctn.Resolve("game-usecase").(usecase.GameUsecase)

	game, apiError := useCase.Flag(ID, square.Row, square.Col)
	if apiError != nil {
		c.AbortWithStatusJSON(apiError.Status, apiError.Error())
		return
	}

	c.JSON(http.StatusOK, game)
	return
}
