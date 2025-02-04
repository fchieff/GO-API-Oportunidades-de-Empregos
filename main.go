package main

import (
	"github.com/fchieff/GO-API-Oportunidades-de-Empregos.git/config"
	"github.com/fchieff/GO-API-Oportunidades-de-Empregos.git/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// Initialize the config

	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	// Initialize the router
	router.Initialize()
}
