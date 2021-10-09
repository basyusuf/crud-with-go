package database

import (
	"fmt"
	"log"
	"main/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseConnector *gorm.DB
var err error

const connectionString = "host=localhost user=postgres password=postgres dbname=case_crud port=5432 sslmode=disable TimeZone=Europe/Istanbul"

func Connect() {
	DatabaseConnector, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database! Please check credentials or database status")
	}
	log.Println("Connection was successful!")
	DatabaseConnector.AutoMigrate(&models.User{})
}
