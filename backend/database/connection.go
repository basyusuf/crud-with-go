package database

import (
	"fmt"
	"log"
	"main/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseConnector *gorm.DB
var err error

func Connect() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	connectionString := fmt.Sprintf("postgres://%s/%s?sslmode=disable&user=%s&password=%s&port=%s", dbHost, dbName, username, password, dbPort)
	//connectionString := "postgres://db/case_crud?sslmode=disable&user=postgres&password=postgres"
	DatabaseConnector, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database! Please check credentials or database status")
	}
	log.Println("Connection was successful!")
	DatabaseConnector.AutoMigrate(&models.User{})
}
