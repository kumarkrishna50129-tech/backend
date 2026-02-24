package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name              string     `gorm:"size:100;not null" json:"name"`
	Email             string     `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password          string     `gorm:"not null" json:"-"` // JSON mein password nahi jayega
	IsPremium         bool       `gorm:"default:false" json:"is_premium"`
	SubscriptionUntil *time.Time `json:"subscription_until"`
	Results           []Result   `gorm:"foreignKey:UserID" json:"results,omitempty"`
}
