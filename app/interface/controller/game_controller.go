package controller

import (
	"net/http"

	"github.com/egorkos/minesweeper/app/domain/model"
	"github.com/egorkos/minesweeper/app/registry"
	"github.com/egorkos/minesweeper/app/usecase"
	"github.com/gin-gonic/gin"
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
