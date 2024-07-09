package database

import (
	"log"
	"note/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	config.LoadConfig()
	dsn := config.GetDBConnectionString()
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//database.AutoMigrate(&model.User{}, &model.Todo{})

	DB = database
}
