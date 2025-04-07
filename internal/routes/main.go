package routes

import (
	"contents-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine, vodController controllers.VodContentController, imgController controllers.ImageController, distributorController controllers.ContentDistributorController, categoryController controllers.CategoryController) {
	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"response": "Server is Running"})
		})

		// Rotas de conteúdos
		api.GET("/contents", vodController.GetAllVodContents)
		api.GET("/content/:id", vodController.GetVodContentByID)
		api.POST("/content", vodController.InsertVodContent)

		// Rotas de imagens
		api.GET("/images", imgController.FindAll)
		api.GET("/image/:id", imgController.FindByID)
		api.POST("/upload", imgController.InsertImage)

		api.GET("/distributors", distributorController.FindAll)
		api.GET("/distributor/:id", distributorController.FindByID)
		api.POST("/distributor", distributorController.Create)

		api.GET("/categories", categoryController.FindAll)
		api.GET("/category/:id", categoryController.FindByID)
	}
}
