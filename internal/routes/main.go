package routes

import (
	"contents-api/internal/app"
	"contents-api/internal/controllers"
	categoriesRoutes "contents-api/internal/routes/categories_routes"
	channelsRoutes "contents-api/internal/routes/channels_routes"
	distributorsRoutes "contents-api/internal/routes/distributors_routes"
	imgsRoutes "contents-api/internal/routes/images_routes"
	vodsRoutes "contents-api/internal/routes/vods_routes"

	"github.com/gin-gonic/gin"
)

var vodContentController controllers.VodContentController
var channelsController controllers.ChannelsController
var imageController controllers.ImageController
var distributorController controllers.ContentDistributorController
var categoryController controllers.CategoryController

func LoadControllers(deps *app.DependencyContainer) {
	vodContentController = deps.VodContentController
	channelsController = deps.ChannelsController
	imageController = deps.ImageController
	distributorController = deps.ContentDistributorController
	categoryController = deps.CategoryController
}

func AppRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"response": "Server is Running"})
		})

		vodsRoutes.SetupVodsRoutes(api, &vodContentController)

		channelsRoutes.SetupChannelsRoutes(api, &channelsController)

		imgsRoutes.SetupImagesRoutes(api, &imageController)

		distributorsRoutes.SetupDistributorsRoutes(api, &distributorController)

		categoriesRoutes.SetupCategoriesRoutes(api, &categoryController)
	}
}
