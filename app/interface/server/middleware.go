package server

import (
	"github.com/Sirupsen/logrus"
	"github.com/egorkos/minesweeper/app/registry"
	"github.com/gin-gonic/gin"
)

func InjectContainer() gin.HandlerFunc {
	logrus.Debug("Starting container")
	ctn, err := registry.NewContainer()
	if err != nil {
		logrus.Fatalf("failed to build container: %v", err)
	}
	return func(c *gin.Context) {
		c.Set("ctn", ctn)
		c.Next()
	}
}
