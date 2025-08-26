package main

import (
	"hello-gin/config"
	"hello-gin/internal/migrations"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Check command line arguments
	if len(os.Args) < 2 {
		log.Println("Usage: go run cmd/migrate/main.go [migrate|drop|reset]")
		log.Println("  migrate - Run migrations")
		log.Println("  drop    - Drop all tables")
		log.Println("  reset   - Drop all tables and run migrations")
		os.Exit(1)
	}

	command := os.Args[1]

	// Connect to database
	config.ConnectDB()

	switch command {
	case "migrate":
		if err := migrations.RunMigrations(config.DB); err != nil {
			log.Fatal("Migration failed:", err)
		}
		log.Println("✅ Migration completed!")

	case "drop":
		if err := migrations.DropAllTables(config.DB); err != nil {
			log.Fatal("Drop tables failed:", err)
		}
		log.Println("✅ All tables dropped!")

	case "reset":
		if err := migrations.DropAllTables(config.DB); err != nil {
			log.Fatal("Drop tables failed:", err)
		}
		if err := migrations.RunMigrations(config.DB); err != nil {
			log.Fatal("Migration failed:", err)
		}
		log.Println("✅ Database reset completed!")

	default:
		log.Printf("Unknown command: %s\n", command)
		log.Println("Available commands: migrate, drop, reset")
		os.Exit(1)
	}
}
