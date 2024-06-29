package main

import (
	"api-openings/config"
	"api-openings/router"
)

var (
	logger config.Logger
)
func main() {

	logger = *config.NewLogger("main")

	// Initialize configs
	error := config.InitDB()
	if error != nil {
		logger.Err("Failed to initialize database:", error)
  }
	// Initialize the router
	router.Initialize()
}
