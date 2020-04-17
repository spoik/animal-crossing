package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spoik/animal-crossing/admin"
	"github.com/spoik/animal-crossing/handlers"
	"net/http"
)

func Create(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	mux := http.NewServeMux()
	Admin := admin.Setup(db)

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.GET("/items", handlers.AllBooks)

	Admin.MountTo("/admin", mux)
	router.Any("/admin/*resources", gin.WrapH(mux))

	return router
}

