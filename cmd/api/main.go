package main

import (
	"log"
	"net/http"
	"os"

	config "kumarkrishna50129-tech/backend/configs"
	"kumarkrishna50129-tech/backend/internal/models"
	"kumarkrishna50129-tech/backend/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. .env file load karein (Error check ke saath)
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// 2. Database Connect karein (Postgres)
	config.ConnectDB()

	// 3. Auto-Migrate Tables
	// Isse Postgres mein automatic Users aur Results table ban jayenge
	err := config.DB.AutoMigrate(&models.User{}, &models.Result{})
	if err != nil {
		log.Fatal("❌ Database Migration Failed:", err)
	}
	log.Println("✅ Database Migration Successful!")

	// 4. Gin Router initialize karein
	r := gin.Default()

	// 5. CORS Middleware Setup (Vite Frontend se connect karne ke liye)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://protypist.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// 6. Health Check Route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "active",
			"engine":  "ProTypist Golang v1.0",
			"revenue": "On track for ₹1 Crore",
		})
	})

	// 7. API Routes Register karein
	routes.SetupRoutes(r)

	// 8. Server Start
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server is flying on http://localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
