package controllers

import (
	channels_service "contents-api/internal/services/channels_services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChannelsController struct {
	service channels_service.Service
}

func NewChannelController(service channels_service.Service) *ChannelsController {
	return &ChannelsController{
		service: service,
	}
}

func (ccc *ChannelsController) FindAll(ctx *gin.Context) {
	channels, err := ccc.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, channels)
}
