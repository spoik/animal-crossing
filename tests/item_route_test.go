package tests

import (
	"encoding/json"
	"testing"

	"github.com/spoik/animal-crossing/models"
	"github.com/spoik/animal-crossing/serializers"
	"github.com/stretchr/testify/assert"
)

func TestNoItemsInDB(t *testing.T) {
	router, _ := Setup()
	response := MakeRequest(router, "GET", "/items")
	assert.Equal(t, 200, response.Code)
	assert.JSONEq(t, `[]`, response.Body.String())
}

func TestItemsInDB(t *testing.T) {
	router, db := Setup()

	item1 := models.Item{ Title: "testItem", Bells: 100 }
	db.Create(&item1)

	item2 := models.Item{ Title: "testItem", Bells: 100 }
	db.Create(&item2)

	response := MakeRequest(router, "GET", "/items")

	assert.Equal(t, 200, response.Code)

	expectedJson, err := json.Marshal(serializers.SerializeItemModels([]models.Item{item1, item2}))
	if err != nil {
		t.Errorf("Unable to ")
	}
	assert.JSONEq(t, string(expectedJson), response.Body.String())
}
