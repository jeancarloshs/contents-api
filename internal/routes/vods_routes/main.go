package vods_routes

import (
	"contents-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupVodsRoutes(router *gin.RouterGroup, controller *controllers.VodContentController) {

	// Rotas de conteúdos
	router.GET("/vods", controller.FindAll)
	router.GET("/vod/:id", controller.FindByID)
	router.POST("/vod", controller.Create)
}
