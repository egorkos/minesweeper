package server

import (
	"net/http"

	"github.com/egorkos/minesweeper/app/interface/controller"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	var router = gin.Default()
	initializeRoutes(router)
	return router
}

func initializeRoutes(router *gin.Engine) {
	router.Use(InjectContainer())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/game", controller.CreateGame)
	router.GET("/games/:id", controller.GetGame)
	router.GET("/games", controller.ListGames)
	router.POST("/games/:id/reveal", controller.Reveal)
	router.POST("/games/:id/flag", controller.Flag)
}
