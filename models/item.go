package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Title string `gorm:"not null" validate:"required"`
	Bells uint   `gorm:"not null" validate:"required"`
}
