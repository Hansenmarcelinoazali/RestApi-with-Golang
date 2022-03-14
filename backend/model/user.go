package model

import (
	"time"
	// "github.com/jinzhu/gorm"
)

type User struct {
	ID        int       `json:"name"`
	Name      string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// type UpdateUser struct {
// 	ID        int       `json:"name"`
// 	Name      string    `json:"username"`

// }
