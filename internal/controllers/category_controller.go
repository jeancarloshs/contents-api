package controllers

import (
	"contents-api/internal/services/category_services"
	"net/http"

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
