package database

import (
	"gorm.io/gorm"
	_ "gorm.io/driver/sqlite"
)

var DB *gorm.DB