package handler

import (
	"api-openings/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil{
		logger.Err("Error creating opening")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro em nossos servidores",
			"status": http.StatusInternalServerError,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Vaga criada com sucesso",
    "status": http.StatusCreated,
    "opening": opening,
	})
}

func ShowOpening(ctx *gin.Context) {}

func DeleteOpening(ctx *gin.Context) {}

func UpdateOpening(ctx *gin.Context) {}

func ShowAllOpenings(ctx *gin.Context) {}
