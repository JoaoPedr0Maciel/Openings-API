package handler

import (
	"api-openings/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1

// GetAllOpenings Delete openings.
// @Summary Delete specific opening
// @Description Delete opening by id
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {string} opening successfully deleted
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/openings/{id} [delete]
func DeleteOpening(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "query parameter (id) is required",
		})
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":  "opening not found",
			"status": http.StatusNotFound,
		})
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":  "error deleting opening",
			"status": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "opening successfully deleted",
		"status":  http.StatusOK,
	})

}