package serializers

import (
	"github.com/spoik/animal-crossing/models"
	"time"
)

type Item struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Bells     uint      `json:"bells"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func SerializeItemModel(item *models.Item) *Item {
	return &Item{
		ID:        item.ID,
		Title:     item.Title,
		Bells:     item.Bells,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func SerializeItemModels(items []models.Item) []*Item {
	newItems := []*Item{}

	for _, item := range items {
		newItems = append(newItems, SerializeItemModel(&item))
	}

	return newItems
}
