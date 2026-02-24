package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// 1. .env file ko load karne ki koshish karein
	_ = godotenv.Load()

	// 2. Render ki DATABASE_URL uthayein
	dsn := os.Getenv("DATABASE_URL")

	// Agar DATABASE_URL khali hai, tabhi localhost use karein
	if dsn == "" {
		fmt.Println("⚠️ DATABASE_URL nahi mili, local connect kar rahe hain...")
		dsn = "host=localhost user=postgres password=YOUR_LOCAL_PASSWORD dbname=protypist_db port=5432 sslmode=disable"
	}

	// 3. Connection establish karein
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Agar yahan error aata hai, toh check karein password ya URL sahi hai ya nahi
		log.Fatal("❌ Database Connection Error: ", err)
	}

	DB = db
	fmt.Println("✅ Database Connected Successfully!")
}
