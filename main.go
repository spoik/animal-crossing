package main

import (
	"github.com/spoik/animal-crossing/handlers"
	"github.com/spoik/animal-crossing/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := models.Setup()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/items", handlers.AllBooks)

	r.Run()
}
