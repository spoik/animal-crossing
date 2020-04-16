package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spoik/animal-crossing/models"
	"github.com/spoik/animal-crossing/router"
	"net/http"
	"net/http/httptest"
	"os"
)

const dbFile = "test.db"

func Setup() (*gin.Engine, *gorm.DB) {
	err := os.Remove(dbFile)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("Unable to delete test database: %s", err))
	}

	db := models.Setup("sqlite3", dbFile)
	return router.Create(db), db
}

func MakeRequest(router *gin.Engine, method string, url string) *httptest.ResponseRecorder {
	httpRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, nil)

	if err != nil {
		panic(fmt.Sprintf("Unable to create request for %s %s", method, url))
	}

	router.ServeHTTP(httpRecorder, request)
	return httpRecorder
}

