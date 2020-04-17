package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spoik/animal-crossing/models"
	"github.com/spoik/animal-crossing/router"
	"io"
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

func MakeRequest(router *gin.Engine, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	httpRecorder := httptest.NewRecorder()
	request, err := http.NewRequest(method, url, body)

	if err != nil {
		panic(fmt.Sprintf("Unable to create %s request for %s", method, url))
	}

	router.ServeHTTP(httpRecorder, request)
	return httpRecorder
}

func GetRequest(router *gin.Engine, url string) *httptest.ResponseRecorder {
	return MakeRequest(router, "GET", url, nil)
}

func PostRequest(router *gin.Engine, url string, body io.Reader) *httptest.ResponseRecorder {
	return MakeRequest(router, "POST", url, body)
}
