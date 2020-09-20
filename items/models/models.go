package models

import (
	"gorm.io/gorm"
	_ "gorm.io/driver/sqlite"
)

type Item struct {
	gorm.Model
	Title  string `json:"title"`
	Category string    `json:"category"`
	Description string `json:"description"`
	OwnerID string `json:"ownerId"`
}
