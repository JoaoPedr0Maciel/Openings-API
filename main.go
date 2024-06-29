package main

import (
	"api-openings/config"
	"api-openings/router"
)

func main() {

	// Initialize configs
	error := config.InitDB()
	if error != nil {
    panic(error)
  }
	// Initialize the router
	router.Initialize()
}
