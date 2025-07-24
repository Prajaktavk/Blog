package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // Load .env file
	"gorm.io/driver/mysql"     // MySQL driver
	"gorm.io/gorm"             // GORM ORM
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Fetch DB credentials from env
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Format DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	// Connect to the database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to MySQL:", err)
	}

	// Save to global variable
	DB = database

	log.Println("✅ Connected to MySQL database")
}
