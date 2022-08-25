package repository

import (
	"hello/models"

	"gorm.io/gorm"
)

func Migration(db gorm.DB) {
	db.AutoMigrate(&models.User{})
}
