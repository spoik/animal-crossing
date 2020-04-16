package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Setup(DBDialect string, DBSource string) *gorm.DB {
	db, err := gorm.Open(DBDialect, DBSource)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Item{})

	return db
}
