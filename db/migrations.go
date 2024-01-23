package db

import (
	"github.com/biqilbooks/models"
)

func SyncDatabase() {
	// Migrate the schema
	DB.AutoMigrate(&models.Tweet{})
}
