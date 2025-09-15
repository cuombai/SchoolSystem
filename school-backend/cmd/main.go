package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"schoolsystem/school-backend/api"
	"schoolsystem/school-backend/config"
)

func main() {
	//Load environment variables

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")

	}

	//Connect to MongoDB
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	//initialize router
	router := api.SetupRouter()

	//start server
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}