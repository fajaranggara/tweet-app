package main

import (
	"tweet-app/config"
	"tweet-app/routes"
)

func main() {
	// for load godotenv
	// environment := utils.GetEnv("ENVIRONMENT", "development")

	// if environment == "development" {
	// 	err := godotenv.Load()
  
	// 	if err != nil {
	// 	  log.Fatal("Error loading .env file")
	// 	}
	//   }


	// database connection
	db := config.ConnectDatabase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	//router
	r := routes.SetupRouter(db)

	r.Run()
}