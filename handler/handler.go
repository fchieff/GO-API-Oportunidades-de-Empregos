package handler

import (
	"github.com/fchieff/GO-API-Oportunidades-de-Empregos.git/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
