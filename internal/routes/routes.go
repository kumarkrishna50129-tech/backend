package routes

import (
	"kumarkrishna50129-tech/backend/internal/controllers"
	"kumarkrishna50129-tech/backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes handles all API endpoints
func SetupRoutes(r *gin.Engine) {
	// Base API Group
	api := r.Group("/api")
	{
		// --- PUBLIC ROUTES (No Token Required) ---
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register) // Signup
			auth.POST("/login", controllers.Login)       // Login
		}

		// --- PROTECTED ROUTES (JWT Token Required) ---
		// In routes par jane se pehle Middleware token check karega
		userAction := api.Group("/user")
		userAction.Use(middleware.AuthRequired())
		{
			// Typing Results Management
			userAction.POST("/save-result", controllers.SaveResult) // Result save karna
			userAction.GET("/history", controllers.GetUserHistory)  // User ki history dekhna

			// Profile Management
			// userAction.GET("/profile", controllers.GetProfile)      // Profile details
		}

		// --- PREMIUM ROUTES (Logic for Paywall) ---
		premium := api.Group("/premium")
		premium.Use(middleware.AuthRequired()) // Pehle Login check
		// premium.Use(middleware.PremiumCheck()) // Phir Payment check
		{
			// premium.GET("/exclusive-exams", controllers.GetMockExams)
		}
	}
}
