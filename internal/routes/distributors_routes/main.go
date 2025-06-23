package distributors_routes

import (
	"contents-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupDistributorsRoutes(router *gin.RouterGroup, controller *controllers.ContentDistributorController) {

	router.GET("/distributors", controller.FindAll)
	router.GET("/distributor/:id", controller.FindByID)
	router.POST("/distributor", controller.Create)
}
