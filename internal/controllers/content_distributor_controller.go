package controllers

import (
	"contents-api/internal/services/content_distributor_services"
)

type ContentDistributorController struct {
	services content_distributor_services.Service
}

func NewContentDistributorController(service content_distributor_services.Service) *ContentDistributorController {
	return &ContentDistributorController{
		services: service,
	}
}
