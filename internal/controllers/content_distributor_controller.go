package controllers

import (
	"contents-api/internal/services/content_distributor_services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentDistributorController struct {
	services content_distributor_services.Service
}

func NewContentDistributorController(service content_distributor_services.Service) *ContentDistributorController {
	return &ContentDistributorController{
		services: service,
	}
}

func (gacdc *ContentDistributorController) GetAllContentDistributorController(ctx *gin.Context) {
	getAll, err := gacdc.services.GetAllContentDistributorService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(getAll) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "Nenhuma distribuidora encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, getAll)
}
