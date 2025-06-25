package channels_routes

import (
	"contents-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupChannelsRoutes(router *gin.RouterGroup, controller *controllers.ChannelsController) {

	router.GET("/channels", controller.FindAll)
}
