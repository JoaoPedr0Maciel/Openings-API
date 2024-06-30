package handler

import (
	"api-openings/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1

// GetSpecifOpening Get specific opening
// @Summary Get specific opening
// @Description Get just one opening by id
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {string} OneOpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/openings/{id} [get]
func ShowOpening(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "query parameter (id) is required",
			"status": http.StatusBadRequest,
		})
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err!= nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "opening not found",
			"message": http.StatusNotFound,                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
    "opening": opening,
    "status":  http.StatusOK,
  })
	
}