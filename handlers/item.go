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
