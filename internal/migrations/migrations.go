package migrations

import (
	"fmt"
	"hello-gin/internal/models"
	"log"

	"gorm.io/gorm"
)

// RunMigrations executes all database migrations
func RunMigrations(db *gorm.DB) error {
	log.Println("🔄 Starting database migrations...")

	// Auto migrate all tables
	err := db.AutoMigrate(
		&models.Class{},
		&models.Student{},
		&models.Teacher{},
		&models.Event{},
		&models.AttendanceSession{},
		&models.Attendance{},
	)

	if err != nil {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	log.Println("✅ Database migrations completed successfully!")
	return nil
}

// DropAllTables drops all tables (use with caution!)
func DropAllTables(db *gorm.DB) error {
	log.Println("⚠️  Dropping all tables...")

	err := db.Migrator().DropTable(
		&models.Attendance{},
		&models.AttendanceSession{},
		&models.Event{},
		&models.Student{},
		&models.Teacher{},
		&models.Class{},
	)

	if err != nil {
		return fmt.Errorf("failed to drop tables: %v", err)
	}

	log.Println("✅ All tables dropped successfully!")
	return nil
}
