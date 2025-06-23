package categories_routes

import (
	"contents-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCategoriesRoutes(router *gin.RouterGroup, controller *controllers.CategoryController) {

	router.GET("/categories", controller.FindAll)
	router.GET("/category/:id", controller.FindByID)
}
