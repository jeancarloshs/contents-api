package app

import (
	"contents-api/internal/config"
	"contents-api/internal/controllers"

	repo_cat "contents-api/internal/repository/category_repositories"
	repo_cdr "contents-api/internal/repository/content_distributor_repository"
	repo_img "contents-api/internal/repository/images_repositories"
	repo_vod "contents-api/internal/repository/vod_content_repositories"
	service_cat "contents-api/internal/services/category_services"
	service_cds "contents-api/internal/services/content_distributor_services"
	service_img "contents-api/internal/services/images_services"
	service_vod "contents-api/internal/services/vod_content_services"
)

// DependencyContainer carrega todas as dependências
type DependencyContainer struct {
	VodContentController         controllers.VodContentController
	ImageController              controllers.ImageController
	ContentDistributorController controllers.ContentDistributorController
	CategoryController           controllers.CategoryController
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
	ContentDistributorRepository := repo_cdr.NewContentDistributorRepository(dataBase)
	CategoryRepository := repo_cat.NewCategoryRepository(dataBase)

	// Inicializa os serviços
	VodContentService := service_vod.NewContentVodService(VodContentRepository)
	ImageService := service_img.UploadImageService(&ImageRepository, minioClient)
	ContentDistributorService := service_cds.NewContentVodService(&ContentDistributorRepository)
	CategoryService := service_cat.NewCategoryService(&CategoryRepository)

	// Inicializa os controllers
	VodContentController := controllers.NewVodContentController(VodContentService)
	ImageController := controllers.NewImagesController(ImageService)
	ContentDistributorController := controllers.NewContentDistributorController(ContentDistributorService)
	CategoryController := controllers.NewCategoryController(CategoryService)

	return &DependencyContainer{
		VodContentController:         *VodContentController,
		ImageController:              *ImageController,
		ContentDistributorController: *ContentDistributorController,
		CategoryController:           *CategoryController,
	}, nil
}
