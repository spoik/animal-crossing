package admin

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/spoik/animal-crossing/models"
)

func Setup(db *gorm.DB) *admin.Admin {
	admin := admin.New(&admin.AdminConfig{DB: db})
	admin.AddResource(&models.Item{})
	return admin
}
