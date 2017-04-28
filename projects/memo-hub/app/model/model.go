package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Memo .
type Memo struct {
	ID   uint
	Body string `json:"body"`
}

// DBMigrate .
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Memo{})
	return db
}
