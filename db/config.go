package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
)

func Connect() *gorm.DB {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up database connection
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")
	// dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	// db, err := gorm.Open("postgres", dbURI)
	// if err != nil {
	// 	log.Fatal("Failed to connect to database")
	// }

	dbConString := os.Getenv("DB_CON_STRING")
	db, err := gorm.Open(mysql.Open(dbConString), &gorm.Config{})

	if err != nil {
		fmt.Println("there was an error connecting")
	}
	// Enable logging if enabled in environment
	// debug := os.Getenv("DB_DEBUG")
	// if debug == "true" {
	// 	// db.LogMode(true)
	// }

	return db
}
