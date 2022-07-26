package models

import "time"

type Like struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	PostID    uint      `json:"post_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User	  User		`json:"-"`
	Post	  Post		`json:"-"`
}