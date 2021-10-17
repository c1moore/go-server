package main

import (
	"net/http"

	"github.com/c1moore/go-server/messages"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"healthy": true,
		})
	})

	messages.Init(engine)

	engine.Run()
}
