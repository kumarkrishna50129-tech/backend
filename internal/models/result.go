package models

import (
	"time"

	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	UserID    uint    `json:"user_id"`
	WPM       int     `json:"wpm"`
	Accuracy  float64 `json:"accuracy"`
	Language  string  `json:"language" gorm:"size:20"` // English/Hindi
	ExamMode  bool    `json:"exam_mode" gorm:"default:false"`
	CreatedAt time.Time
}
