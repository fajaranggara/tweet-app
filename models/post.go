package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Content   string    `json:"content" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User	  User		`json:"-"`
	Like	  []Like	`json:"-"`
}