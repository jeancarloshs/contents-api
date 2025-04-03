package images_services

import (
	model "contents-api/internal/models"
	"contents-api/internal/repository/images_repositories"

	"github.com/minio/minio-go/v7"
)

type Service interface {
	GetAllImagesService() ([]model.Images, error)
	GetImageByIDService(imgID int) ([]model.Images, error)
	CreateImageService(imgContent model.Images) (model.Images, error)
}

type ImageService struct {
	repository images_repositories.Repository
	minio      *minio.Client
}

func UploadImageService(imageContent images_repositories.Repository, minioClient *minio.Client) *ImageService {
	return &ImageService{
		repository: imageContent,
		minio:      minioClient,
	}
}
