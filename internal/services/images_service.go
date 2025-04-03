package services

import (
	"fmt"
	model "nx-api/internal/models"
	"nx-api/internal/repository"

	"github.com/minio/minio-go/v7"
)

type ImageService struct {
	repo  repository.ImageRepository
	minio *minio.Client
}

func UploadImageService(imageContent repository.ImageRepository, minioClient *minio.Client) *ImageService {
	return &ImageService{
		repo:  imageContent,
		minio: minioClient,
	}
}

func (gais *ImageService) GetAllImagesService() ([]model.Images, error) {
	return gais.repo.GetAllImageRepository()
}

func (gibis *ImageService) GetImageByIDService(imgID int) (model.Images, error) {
	return gibis.repo.GetImageByIDRepository(imgID)
}

func (uis *ImageService) UploadImage(imageContent model.Images) (model.Images, error) {
	imgInsert, err := uis.repo.UploadImageRepository(imageContent)
	if err != nil {
		fmt.Println(err)
		return model.Images{}, nil
	}

	return imgInsert, nil
}

func (uims *ImageService) UploadImageMinioService(dataMinio model.UploadToMinIO, minioClient *minio.Client) (model.Images, error) {
	uploadService := model.Images{
		Url: dataMinio.ObjectName,
	}

	_, err := uims.UploadImage(uploadService)
	if err != nil {
		fmt.Println("Erro ao fazer upload da imagem:", err)
		return model.Images{}, err
	}

	// Retorna um objeto `model.Images` com a URL do MinIO (ajuste conforme necessário)
	return model.Images{
		Url: dataMinio.ObjectName,
	}, nil
}
