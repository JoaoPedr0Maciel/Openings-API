package handler

import (
	"api-openings/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1

// GetSpecifOpening Update specific opening
// @Summary Update specific opening
// @Description Update an opening by id
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {string} UpdateOpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/openings/{id} [put]
func UpdateOpening(ctx *gin.Context) {
	id := ctx.Query("id")
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
      "status": http.StatusBadRequest,
    })
    return
	}

	if id == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "query parameter (id) is required",
      "status": http.StatusBadRequest,
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

	// Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary != "" {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
      "error":  "error updating opening",
      "status": http.StatusInternalServerError,
    })
    return
	}

	ctx.JSON(http.StatusOK, gin.H{
    "message": "opening successfully updated",
    "status":  http.StatusOK,
    "opening": opening,
  })
}