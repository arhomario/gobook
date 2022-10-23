package models

import (
	"time"
)

type Book struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	Title      string    `json:"title" gorm:"size:255;not null"`
	Content    string    `json:"content" gorm:"size:255;not null"`
	Created_at time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Updated_at time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
