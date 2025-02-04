package config

import (
	"os"

	"github.com/fchieff/GO-API-Oportunidades-de-Empregos.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if the database exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		// Create the database directory if it doesn't exist
		logger.Info("sqlite database not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("failed to create database directory: %v", err)
			return nil, err
		}

		// Create the database file
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("failed to create database file: %v", err)
			return nil, err
		}
		file.Close() // Close the file after creating it
	}

	// Create database connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite open error: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite auto migrate error: %v", err)
		return nil, err
	}

	// Return the database
	return db, nil
}
