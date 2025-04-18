package controllers

import (
	model "contents-api/internal/models"
	vod_content "contents-api/internal/services/vod_content_services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VodContentController struct {
	contentVodService vod_content.Service
}

func NewVodContentController(service vod_content.Service) *VodContentController {
	return &VodContentController{
		contentVodService: service,
	}
}

func (gacvd *VodContentController) FindAll(ctx *gin.Context) {
	allContents, err := gacvd.contentVodService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(allContents) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "Nenhum conteudo encontrado",
		})
		return
	}

	ctx.JSON(http.StatusOK, allContents)
}

func (gacvd *VodContentController) FindByID(ctx *gin.Context) {
	vodID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	allContents, err := gacvd.contentVodService.FindByID(vodID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, allContents)
}

func (ivc *VodContentController) Create(ctx *gin.Context) {
	var vodContent model.VodContent
	err := ctx.BindJSON(&vodContent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de dados inválido"})
		return
	}

	contentInsert, err := ivc.contentVodService.Create(vodContent)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, contentInsert)
}
