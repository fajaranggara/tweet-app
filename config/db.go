package config

import (
	"fmt"
	"tweet-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	username := "root"
	password := "root"
	host := "tcp(127.0.0.1:3306)"
	database := "db_tweetapp"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Like{})

	return db
}