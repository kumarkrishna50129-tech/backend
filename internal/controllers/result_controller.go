package controllers

import (
	config "kumarkrishna50129-tech/backend/configs"
	models "kumarkrishna50129-tech/backend/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ResultRequest defines the structure for incoming typing data
type ResultRequest struct {
	WPM       int     `json:"wpm" binding:"required"`
	Accuracy  float64 `json:"accuracy" binding:"required"`
	Language  string  `json:"language" binding:"required"`
	Mistakes  int     `json:"mistakes"`
	TimeTaken int     `json:"time_taken"` // seconds mein
}

// SaveResult handles POST /api/results
func SaveResult(c *gin.Context) {
	var req ResultRequest

	// 1. JSON Bind aur Validation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	// 2. Anti-Cheat Check (Simple Logic)
	// Agar accuracy 100% hai aur WPM 200 se upar, toh flag karein
	if req.WPM > 200 && req.Accuracy > 99 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Suspicious activity detected!"})
		return
	}

	// 3. User Identity (JWT se UserID nikalna - Abhi hardcoded hai, baad mein Middleware se aayega)
	userID := uint(1)

	// 4. Model taiyar karein
	result := models.Result{
		UserID:    userID,
		WPM:       req.WPM,
		Accuracy:  req.Accuracy,
		Language:  req.Language,
		CreatedAt: time.Now(),
	}

	// 5. Database mein Save karein
	dbResult := config.DB.Create(&result)
	if dbResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save result"})
		return
	}

	// 6. Success Response
	c.JSON(http.StatusOK, gin.H{
		"message": "Result saved successfully",
		"data": gin.H{
			"wpm":      result.WPM,
			"accuracy": result.Accuracy,
			"id":       result.ID,
		},
	})
}

// GetUserHistory handles GET /api/results/history
func GetUserHistory(c *gin.Context) {
	var results []models.Result
	userID := uint(1) // Placeholder

	// Last 10 results fetch karein
	config.DB.Where("user_id = ?", userID).Order("created_at desc").Limit(10).Find(&results)

	c.JSON(http.StatusOK, gin.H{
		"history": results,
	})
}
