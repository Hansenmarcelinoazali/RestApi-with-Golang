package model

import (
	"time"
	// "github.com/jinzhu/gorm"
)

type Note struct {
	ID          int       `json:"id"`
	UserId      int       `json:"userid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
