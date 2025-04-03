package controllers

import (
	"contents-api/internal/config"
	"contents-api/internal/services/images_services"
	services "contents-api/internal/services/minio_services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	imageService images_services.Service
	minioService *services.MinIOUploadService
}

func (gaic *ImageController) GetAllImagesController(ctx *gin.Context) {
	allImages, err := gaic.imageService.GetAllImagesService()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(allImages) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "Nenhuma imagem encontrada",
		})
		return
	}

	ctx.JSON(http.StatusOK, allImages)
}

func (gibic *ImageController) GetImageByIDController(ctx *gin.Context) {
	imgID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	getImg, err := gibic.imageService.GetImageByIDService(imgID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, getImg)
}

// Modifique o construtor para inicializar o MinIOUploadService
func NewImagesController(service images_services.Service) *ImageController {
	// Obter cliente MinIO da configuração global
	minioClient, err := config.ConnectMinio()
	if err != nil {
		log.Fatalf("Falha ao conectar ao MinIO: %v", err)
	}

	// Criar serviço de upload
	minioService := services.NewMinIOUploadService(minioClient, "bucket-images")

	return &ImageController{
		imageService: service,
		minioService: minioService, // Agora está sendo inicializado
	}
}

func (im *ImageController) InsertImage(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao obter arquivo"})
		log.Println("Erro ao obter arquivo:", err)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao abrir arquivo"})
		log.Println("Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	uploadOptions := services.UploadOptions{
		ObjectName:  fileHeader.Filename,
		ContentType: fileHeader.Header.Get("Content-Type"),
	}

	uploadInfo, err := im.minioService.UploadFile(ctx, file, fileHeader.Size, uploadOptions)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ // Mudei para 500 (erro interno)
			"error":   "Erro ao salvar arquivo",
			"details": err.Error(),
		})
		log.Println("Erro ao salvar arquivo:", err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":  "Arquivo carregado com sucesso",
		"name":     fileHeader.Filename,
		"size":     fileHeader.Size,
		"location": uploadInfo.Bucket + "/" + uploadInfo.Key,
	})
}
