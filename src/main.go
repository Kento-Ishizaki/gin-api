package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go_test/controller"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	engine.GET("/posts", controllers.Getting)

	engine.Run(":3000")
}
