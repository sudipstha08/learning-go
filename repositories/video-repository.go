package repository

import (
	"learning-go/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type VideoRepository interface {
	Save(video models.Video)
	Update(video models.Video)
	Delete(video models.Video)
	FindAll() []models.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&models.Video{}, &models.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(video models.Video) {
	// reference to the struct
	db.connection.Create(&video)
}
func (db *database) Update(video models.Video) {
	db.connection.Save(&video)
}
func (db *database) Delete(video models.Video) {
	db.connection.Delete(&video)
}
func (db *database) FindAll() []models.Video {
	var videos []models.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
