package config

import (
	"gorm.io/gorm"
)

//database

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(prefix string) *Logger {
	//initialize logger
	logger = NewLogger(prefix)
	return logger
}
