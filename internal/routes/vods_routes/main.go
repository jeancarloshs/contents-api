package vod_routes

import (
	"contents-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupVodsRoutes(router *gin.Engine) {
	controller := &controllers.VodContentController{}

	// Rotas de conteúdos
	router.GET("/contents", controller.FindAll)
	router.GET("/content/:id", controller.FindByID)
	router.POST("/content", controller.Create)
}
