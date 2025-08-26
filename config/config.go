package config

import (
	"fmt"
	"hello-gin/internal/migrations"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}
}

func ConnectDB() {
	// Use environment variables with fallback to default values
	host := getEnvWithDefault("DB_HOST", "localhost")
	user := getEnvWithDefault("DB_USER", "postgres")
	password := os.Getenv("DB_PASSWORD") // Bắt buộc phải có từ .env
	dbname := os.Getenv("DB_NAME")       // Bắt buộc phải có từ .env
	port := getEnvWithDefault("DB_PORT", "5432")

	// Kiểm tra các biến bắt buộc
	if password == "" {
		log.Fatal("❌ DB_PASSWORD is required in .env file")
	}
	if dbname == "" {
		log.Fatal("❌ DB_NAME is required in .env file")
	}

	// First, connect to PostgreSQL without specifying a database to create the database if needed
	adminDSN := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		host, user, password, port,
	)

	log.Println("🔗 Kết nối PostgreSQL để tạo database...")
	adminDB, err := gorm.Open(postgres.Open(adminDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Không thể kết nối PostgreSQL: ", err)
	}

	// Create database if it doesn't exist
	createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbname)
	result := adminDB.Exec(createDBQuery)
	if result.Error != nil {
		// Database might already exist, which is fine
		log.Printf("ℹ️ Database '%s' có thể đã tồn tại: %v", dbname, result.Error)
	} else {
		log.Printf("✅ Đã tạo database '%s' thành công!", dbname)
	}

	// Now connect to the specific database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		host, user, password, dbname, port,
	)

	log.Printf("🔗 Kết nối database '%s'...", dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Kết nối DB thất bại: ", err)
	}

	log.Println("✅ Kết nối PostgreSQL thành công!")
	DB = db

	// Run migrations
	log.Println("🚀 Bắt đầu migration...")
	err = migrations.RunMigrations(db)
	if err != nil {
		log.Fatal("❌ Migration thất bại: ", err)
	}
}

// Helper function to get environment variable with default value
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
