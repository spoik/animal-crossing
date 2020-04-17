package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spoik/animal-crossing/models"
	"github.com/spoik/animal-crossing/serializers"
	"net/http"
)

func AllBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var items []models.Item
	db.Find(&items)

	c.JSON(http.StatusOK, serializers.SerializeItemModels(items))
}

type NewItem struct {
	Title string `json:"title"`
	Bells uint   `json:"bells"`
}

func CreateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var newBook NewItem
	c.ShouldBindJSON(&newBook)

	item := models.Item{Title: newBook.Title, Bells: newBook.Bells}
	db.Create(&item)

	c.Status(http.StatusCreated)
}
