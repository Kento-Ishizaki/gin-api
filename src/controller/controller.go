package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "OK",
	})
}
