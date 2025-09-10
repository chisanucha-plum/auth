package database

import (
	"auth/internal/configuration"
	"auth/internal/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		configuration.Envs.DBHost,
		configuration.Envs.DBUser,
		configuration.Envs.DBPassword,
		configuration.Envs.DBName,
		configuration.Envs.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return db, nil
}

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		configuration.Envs.DBHost,
		configuration.Envs.DBUser,
		configuration.Envs.DBPassword,
		configuration.Envs.DBName,
		configuration.Envs.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to PostgreSQL: %v", err)
		// Try connecting to PostgreSQL first.
		// Fallback to SQLite (may not work if CGO disabled)
		db, err = gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to any database:", err)
		}
		log.Println("Connected to SQLite database: auth.db")
	} else {
		log.Println("Connected to PostgreSQL database")
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully")
	return db
}
