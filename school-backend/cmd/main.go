package main

import (
	"log"
	"os"
	"schoolsystem/school-backend/api"
	"schoolsystem/school-backend/config"
	// "schoolsystem/school-backend/internal/auth"
	// "schoolsystem/school-backend/models"
	// "schoolsystem/school-backend/repository"
	// "time"

	"github.com/joho/godotenv"
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

	// //setup test email for admin
	// hashed, err := auth.HashPassword("adminpassword")
	// if err != nil {
	// 	log.Fatalf("Password hash failed: %v", err)
	// }

	// admin := models.User{
	// 	Name: "Admin",
	// 	Email: "admin@example.com",
	// 	Password: hashed,
	// 	Role: "admin",
	// 	Phone: "0746881441",
	// 	CreatedAt: time.Now(),
	// }

	// if err := repository.InsertUser(admin); err != nil {
	// 	log.Fatalf("Insert failed: %v", err)
	// }
	// log.Println("Admin user created succesfuly")


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