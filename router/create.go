package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spoik/animal-crossing/handlers"
)

func Create(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.GET("/items", handlers.AllBooks)

	return router
}

