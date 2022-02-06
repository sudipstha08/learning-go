package infrastructures

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	DB *gorm.DB
}

func ConnectDatabase() Database {
	// Creating a new connection
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database !")
	}

	// Migrate the database schema using AutoMigrate
	// Call this method on each model created
	// db.AutoMigrate(&models.Book{})

	// Populate DB variable with our database instance
	return Database{
		DB: db,
	}
}
