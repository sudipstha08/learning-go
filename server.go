package main

import (
	"io"
	repository "learning-go/api/repositories"
	"learning-go/api/routes"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
)

// Creates a log file
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	Router := gin.Default()

	// CLOSE DATABASE
	defer videoRepository.CloseDB()

	setupLogOutput()

	// LOAD ENV VARIABLES
	godotenv.Load(".env")

	// INIT ROUTES
	routes.GetRoutes()

	// RUN APP
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	Router.Run(":" + port)
}
