package images_routes

import (
	"contents-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupImagesRoutes(router *gin.Engine) {
	controller := &controllers.ImageController{}

	// Rotas de imagens
	router.GET("/images", controller.FindAll)
	router.GET("/image/:id", controller.FindByID)
	router.POST("/upload", controller.InsertImage)
}
