package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/milsoncodes/gin-template/models"
)

func main() {
	fmt.Println("Hello, Gin!")

	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "Hello, API!",
		})
	})
	r.Run()
}
