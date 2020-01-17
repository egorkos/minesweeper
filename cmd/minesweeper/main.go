package main

import (
	"net/http"

	"github.com/egorkos/minesweeper/app/interface/controller"
	"github.com/egorkos/minesweeper/app/interface/server"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	initializeRoutes()
	router.Run()
}

func initializeRoutes() {
	router.Use(server.InjectContainer())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/game", controller.CreateGame)
	//router.GET("/games/:id", controller.GetGame)
	//router.GET("/games", controller.ListGames)
	//router.POST("/games/:id/reveal", controller.Reveal)
	//router.POST("/games/:id/flag", controller.Flag)
}
