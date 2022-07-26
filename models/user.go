package models

import (
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        	uint		`json:"id" gorm:"primary_key"`
	Name      	string		`json:"name" gorm:"not null"`
	Email     	string		`json:"email" gorm:"not null;unique"`
	Password  	string		`json:"password" gorm:"not null"`
	Image	  	string		`json:"image"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	Post		[]Post		`json:"-"`
	Like		[]Like		`json:"-"`
}

func (usr *User) SaveUser(db *gorm.DB) (*User, error) {
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &User{}, errPassword
	}

	usr.Password = string(hashedPassword)
	usr.Email = html.EscapeString(strings.TrimSpace(usr.Email))

	var err error = db.Create(&usr).Error
	if err != nil {
		return &User{}, err
	}

	return usr, nil
}

