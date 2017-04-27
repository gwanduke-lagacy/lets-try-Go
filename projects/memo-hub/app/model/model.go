package model

import (
	"github.com/jinzhu/gorm"
)

// Memo .
type Memo struct {
	gorm.Model
	ID   int64  `gorm:"unique" json:"id"`
	Body string `json:"body"`
}

// DBMigrate .
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Memo{})
	return db
}
