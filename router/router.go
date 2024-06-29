package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize router
	router := gin.Default()

	// initialize routes
	initializeRoutes(router)

	// running the server
	router.Run(":8080")
}
