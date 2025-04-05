package controllers

import (
	"contents-api/internal/services/category_services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service category_services.Service
}

func NewCategoryController(service category_services.Service) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

func (gacc *CategoryController) GetAllCategoryController(ctx *gin.Context) {
	categories, err := gacc.service.GetAllCategoryService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(categories) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "Nenhuma categoria encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (gcbic *CategoryController) GetCategoryByIDController(ctx *gin.Context) {
	catID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	categoryID, err := gcbic.service.GetCategoryByIDService(catID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, categoryID)
}
