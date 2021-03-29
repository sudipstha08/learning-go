package main

import (
	"io"
	"learning-go/infrastructures"
	repository "learning-go/repositories"
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
	// CLOSE DATABASE
	defer videoRepository.CloseDB()

	setupLogOutput()

	// LOAD ENV VARIABLES
	godotenv.Load(".env")
	infrastructures.GetRoutes()
}
