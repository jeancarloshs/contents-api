package vods_routes

import (
	"contents-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupVodsRoutes(router *gin.RouterGroup, controller *controllers.VodContentController) {

	// Rotas de conteúdos
	router.GET("/contents", controller.FindAll)
	router.GET("/content/:id", controller.FindByID)
	router.POST("/content", controller.Create)
}
