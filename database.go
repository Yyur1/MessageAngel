package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type MessageRecord struct {
	gorm.Model
	GroupId   string `json:"group_id"`
	UserId    string `json:"user_id"`
	Message   string `json:"message"`
	MessageId string `json:"message_id"`
}

var db *gorm.DB

func InitializeDatabase() {
	ConnectToDatabase()
	CreateMessageRecord()
}

func ConnectToDatabase() {
	user := os.Getenv("POSTGRESQL_USER")
	password := os.Getenv("POSTGRESQL_PASSWORD")
	port := os.Getenv("POSTGRESQL_PORT")
	dsn := fmt.Sprintf("host=localhost user=%v password= dbname=%v port=%v sslmode=disable TimeZone=Asia/Tokyo", user, password, port)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func CreateMessageRecord() {
	err := db.AutoMigrate(&MessageRecord{})
	if err != nil {
		logger.Fatalf("Failed to migrate database: %v", err)
	}
}
