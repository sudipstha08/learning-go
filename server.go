package main

import (
	"fmt"
	"io"
	repository "learning-go/api/repositories"
	"learning-go/api/routes"
	"learning-go/infrastructures"
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

	// Sentry initialization
	infrastructures.InitSentry(Router)

	// CLOSE DATABASE
	defer videoRepository.CloseDB()

	setupLogOutput()

	// LOAD ENV VARIABLES
	godotenv.Load(".env")

	// INIT ROUTES
	routes.GetRoutes(Router)
	fmt.Println("-------- STARTING APPLICATION 🚚 --------")
	fmt.Println("----------------------------------")
	fmt.Println("--------🌱 LEARNING GO 🌱---------")
	fmt.Println("----------------------------------")
	// RUN APP

	port := os.Getenv("PORT")
	if port == "" {
		port = "5005"
	}
	fmt.Printf("------ APPLICATION UP & RUNNING AT PORT %s 🏃 ------", port)
	Router.Run(":" + port)
}
