package controllers

import (
	model "contents-api/internal/models"
	"contents-api/internal/services/content_distributor_services"
	"fmt"
	"net/http"
	"strconv"

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

func (gcbic *ContentDistributorController) GetContentDistributorController(ctx *gin.Context) {
	cdID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	allDistributor, err := gcbic.services.GetByIDContentDistributorService(cdID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, allDistributor)
}

func (ccdc *ContentDistributorController) CreateContentDistributorController(ctx *gin.Context) {
	var distributorContent model.ContentDistributor
	fmt.Println(distributorContent)
	err := ctx.BindJSON(&distributorContent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de dados inválido"})
		return
	}

	distributorInsert, err := ccdc.services.CreateContentDistributorService(distributorContent)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, distributorInsert)
}
