package router

import (
	"api-openings/handler"

	"github.com/gin-gonic/gin"
)

// function to initialize the group of routes
func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("v1")
	{

		// Show openings
		v1.GET("/opening", handler.ShowOpening)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpening)

		v1.PUT("/opening", handler.UpdateOpening)

		v1.GET("/all", handler.ShowAllOpenings)

	}
}
