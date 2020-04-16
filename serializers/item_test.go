package serializers_test

import (
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/jinzhu/gorm"
	"github.com/spoik/animal-crossing/models"
	"github.com/spoik/animal-crossing/serializers"
	"github.com/spoik/animal-crossing/tests"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSingleItem(t *testing.T) {
	item := models.Item{
		Model: gorm.Model{
			ID:        uint(randomdata.Number(0,10)),
			CreatedAt: tests.RandomTime(),
			UpdatedAt: tests.RandomTime(),
			DeletedAt: nil,
		},
		Title: randomdata.FullName(randomdata.Male),
		Bells: uint(randomdata.Number(10,500)),
	}

	itemJson, err := json.Marshal(serializers.SerializeItemModel(&item))

	if err != nil {
		t.Errorf("Unable to marshal JSON for: %+v", item)
	}

	expectedJson := fmt.Sprintf(
		`{
			"bells": %d,
			"created_at": "%s",
			"id": %d,
			"title": "%s",
			"updated_at": "%s"
		}`,
		item.Bells,
		item.CreatedAt.Format(time.RFC3339),
		item.ID,
		item.Title,
		item.UpdatedAt.Format(time.RFC3339),
	)

	assert.JSONEq(t, expectedJson, string(itemJson))
}