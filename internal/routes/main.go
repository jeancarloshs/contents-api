package routes

import (
	categoriesRoutes "contents-api/internal/routes/categories_routes"
	distributorsRoutes "contents-api/internal/routes/distributors_routes"
	imgsRoutes "contents-api/internal/routes/images_routes"
	vodsRoutes "contents-api/internal/routes/vods_routes"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"response": "Server is Running"})
		})

		vodsRoutes.SetupVodsRoutes(router)

		imgsRoutes.SetupImagesRoutes(router)

		distributorsRoutes.SetupDistributorsRoutes(router)

		categoriesRoutes.SetupCategoriesRoutes(router)
	}
}
