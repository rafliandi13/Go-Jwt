package configs

import (
	"GO-JWT/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func ConnectDB(){
	dsn := "root:@tcp(127.0.0.1:3306)/go-jwt"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	DB = db

	log.Println("Database connected")
}