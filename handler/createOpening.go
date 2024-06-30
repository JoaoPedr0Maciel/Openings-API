package handler

import (
	"api-openings/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /v1

// CreateOpeningHandler Creates a new opening.
// @Summary Create an opening
// @Description Create a new open job with the given data
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /v1/openings [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Erro em nossos servidores",
			"status": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Vaga criada com sucesso",
		"status":  http.StatusCreated,
		"opening": opening,
	})
}