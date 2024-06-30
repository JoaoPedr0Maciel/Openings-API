package router

import (
	"api-openings/docs"
	"api-openings/handler"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	// Initialize Swagger documentation
	docs.SwaggerInfo.Title = "Openings API"
	docs.SwaggerInfo.Description = "API for Openings jobs that you can create, list, update and delete"
	docs.SwaggerInfo.Version = "1.0"

	// Configure routes
	v1 := router.Group("/v1")
	{
		v1.GET("/openings/:id", handler.ShowOpening)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings/:id", handler.DeleteOpening)
		v1.PUT("/openings/:id", handler.UpdateOpening)
		v1.GET("/openings", handler.ShowAllOpenings)
	}

	// Serve Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
