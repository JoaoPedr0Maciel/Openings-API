package handler

import (
	"api-openings/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1

// GetAllOpenings Get all openings.
// @Summary Get all openings
// @Description Get the list of openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} GetAllOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/openings [get]
func ShowAllOpenings(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":  "error fetching openings",
			"status": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"openings": openings,
	})
}