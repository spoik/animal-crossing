package tests

import (
	"bytes"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/spoik/animal-crossing/handlers"
	"net/http"
	"strings"
	"testing"

	"github.com/spoik/animal-crossing/models"
	"github.com/spoik/animal-crossing/serializers"
	"github.com/stretchr/testify/assert"
)

func TestNoItemsInDB(t *testing.T) {
	router, _ := Setup()
	response := GetRequest(router, "/items")
	assert.Equal(t, 200, response.Code)
	assert.JSONEq(t, `[]`, response.Body.String())
}

func TestItemsInDB(t *testing.T) {
	router, db := Setup()

	item1 := models.Item{Title: "testItem", Bells: 100}
	db.Create(&item1)

	item2 := models.Item{Title: "testItem", Bells: 100}
	db.Create(&item2)

	response := GetRequest(router, "/items")

	assert.Equal(t, http.StatusOK, response.Code)

	expectedJson, err := json.Marshal(serializers.SerializeItemModels([]models.Item{item1, item2}))
	if err != nil {
		t.Errorf("Unable to ")
	}
	assert.JSONEq(t, string(expectedJson), response.Body.String())
}

func TestSuccessfulCreateItem(t *testing.T) {
	router, db := Setup()

	assertNoItemsInDB(t, db)

	newItem := handlers.NewItem{Title: "testing", Bells: 10}
	postJson, err := json.Marshal(newItem)
	if err != nil {
		panic("Unable to marshal NewItem")
	}

	response := PostRequest(router, "/items", bytes.NewReader(postJson))
	assert.Equal(t, http.StatusCreated, response.Code)

	assertNumItemsInDb(t, db, 1)

	var item models.Item
	db.Model(&models.Item{}).Last(&item)
	assert.Equal(t, item.Title, newItem.Title)
	assert.Equal(t, item.Bells, newItem.Bells)
}

func TestCreateItemWithNoTitle(t *testing.T) {
	router, db := Setup()
	assertNoItemsInDB(t, db)

	requestJson := `{"title": "", "bells": 100}`
	response := PostRequest(router, "/items", strings.NewReader(requestJson))
	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)

	expectedJson := `{"error_messages": ["Title is a required field"]}`
	assert.JSONEq(t, expectedJson, response.Body.String())

	assertNoItemsInDB(t, db)
}

func TestCreateItemWithNoBells(t *testing.T) {
	router, db := Setup()
	assertNoItemsInDB(t, db)

	requestJson := `{"title": "Testing", "bells": 0}`
	response := PostRequest(router, "/items", strings.NewReader(requestJson))
	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)

	expectedJson := `{"error_messages": ["Bells is a required field"]}`
	assert.JSONEq(t, expectedJson, response.Body.String())

	assertNoItemsInDB(t, db)
}

func assertNumItemsInDb(t *testing.T, db *gorm.DB, num int) {
	var count int
	db.Model(&models.Item{}).Count(&count)
	assert.Equal(t, count, num)
}

func assertNoItemsInDB(t *testing.T, db *gorm.DB) {
	assertNumItemsInDb(t, db, 0)
}
