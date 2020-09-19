package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Item struct {
	gorm.Model
	Title  string `json:"name"`
	Owner string `json:"author"`
	Type string    `json:"rating"`
}
