package main

import (
	"github.com/spoik/animal-crossing/models"
	"github.com/spoik/animal-crossing/router"
)

func main() {
	db := models.Setup("sqlite3", "production.db")
	r := router.Create(db)
	r.Run()
}
