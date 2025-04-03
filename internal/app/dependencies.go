package app

import (
	"contents-api/internal/config"
	"contents-api/internal/controllers"
	"contents-api/internal/repository"
	repo_vod "contents-api/internal/repository/vod_content"
	"contents-api/internal/services"
	service_vod "contents-api/internal/services/vod_content"
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
	ImageRepository := repository.GetImageRepository(dataBase)

	// Inicializa os serviços
	VodContentService := service_vod.NewContentVodService(VodContentRepository)
	ImageService := services.UploadImageService(ImageRepository, minioClient)

	// Inicializa os controllers
	VodContentController := controllers.NewVodContentController(VodContentService)
	ImageController := controllers.ImagesController(*ImageService)

	return &DependencyContainer{
		VodContentController: *VodContentController,
		ImageController:      ImageController,
	}, nil
}
