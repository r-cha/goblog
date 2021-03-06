package db

import (
	"r-cha/goblog/config"
	"r-cha/goblog/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var database *gorm.DB
	var err error

	settings := config.NewSettings()
	env := settings.ENVIRONMENT
	if env == "prod" {
		database, err = gorm.Open(postgres.Open(settings.POSTGRES_DSN), &gorm.Config{})
	} else { // if env == "local" {
		database, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect to database!")
	}

	// Register Models
	// TODO: Can this be automated?
	// Tedious thing to remember to do imo
	database.AutoMigrate(&models.Author{})
	database.AutoMigrate(&models.Post{})

	DB = database
}

func Close() {
	conn, _ := DB.DB()
	conn.Close()
}

func Reset() {
	Close()
	Connect()
}
