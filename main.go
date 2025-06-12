package main

import (
	"backend/models"

	"github.com/gin-gonic/gin"
)

// run server with live reload
// go run github.com/pilu/fresh

func main() {
	router := gin.Default()
	models.Migrations()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":9000")
}
