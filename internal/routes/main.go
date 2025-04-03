package routes

import (
	"contents-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine, vodController controllers.VodContentController, imgController controllers.ImageController) {
	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"response": "Server is Running"})
		})

		// Rotas de conteúdos
		api.GET("/contents", vodController.GetAllVodContents)
		api.GET("/content/:id", vodController.GetVodContentByID)
		api.POST("/content/insert", vodController.InsertVodContent)

		// Rotas de imagens
		api.GET("/images", imgController.GetAllImagesController)
		api.GET("/image/:id", imgController.GetImageByIDController)
		api.POST("/upload", imgController.InsertImage)
	}
}
