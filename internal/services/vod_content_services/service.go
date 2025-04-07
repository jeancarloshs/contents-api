package vod_content

import (
	model "contents-api/internal/models"
	vod_content "contents-api/internal/repository/vod_content_repositories"
)

type Service interface {
	FindAll() ([]model.VodContent, error)
	FindByID(vodID int) (model.VodContent, error)
	Create(vodContent model.VodContent) (model.VodContent, error)
}

type VodContentService struct {
	repository vod_content.Repository
}

func NewContentVodService(repo vod_content.Repository) Service {
	return &VodContentService{
		repository: repo,
	}
}
