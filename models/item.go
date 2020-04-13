package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Title string `gorm:"not null"`
	Bells uint   `gorm:"not null"`
}
