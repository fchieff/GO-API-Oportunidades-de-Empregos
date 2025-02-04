package config

import (
	"fmt"

	"gorm.io/gorm"
)

//database

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize SQLite
	db, err = InitSQLite()
	if err != nil {
		return fmt.Errorf("failed to initialize sqlite: %v", err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	//initialize logger
	logger = NewLogger(prefix)
	return logger
}
