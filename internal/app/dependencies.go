package app

import (
	"contents-api/internal/config"
	"contents-api/internal/controllers"
	repo_img "contents-api/internal/repository/images_repositories"
	repo_vod "contents-api/internal/repository/vod_content_repositories"
	service_img "contents-api/internal/services/images_services"
	service_vod "contents-api/internal/services/vod_content_services"
)

// DependencyContainer carrega todas as dependências
type DependencyContainer struct {
	VodContentController controllers.VodContentController
	ImageController      controllers.ImageController
}

// InitializeDependencies carrega todas as dependências da aplicação
func InitializeDependencies() (*DependencyContainer, error) {
	dataBase, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}

	minioClient, err := config.ConnectMinio()
	if err != nil {
		return nil, err
	}

	// Inicializa os repositórios
	VodContentRepository := repo_vod.NewContentVodRepository(dataBase)
	ImageRepository := repo_img.NewImageRepository(dataBase)

	// Inicializa os serviços
	VodContentService := service_vod.NewContentVodService(VodContentRepository)
	ImageService := service_img.UploadImageService(&ImageRepository, minioClient)

	// Inicializa os controllers
	VodContentController := controllers.NewVodContentController(VodContentService)
	ImageController := controllers.NewImagesController(ImageService)

	return &DependencyContainer{
		VodContentController: *VodContentController,
		ImageController:      *ImageController,
	}, nil
}
