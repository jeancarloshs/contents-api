package images_routes

import (
	"contents-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupImagesRoutes(router *gin.RouterGroup, controller *controllers.ImageController) {

	// Rotas de conteúdos
	router.GET("/images", controller.FindAll)
	router.GET("/image/:id", controller.FindByID)
	router.POST("/upload", controller.InsertImage)
}
