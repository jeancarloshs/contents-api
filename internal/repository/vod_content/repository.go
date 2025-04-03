package vod_content

import (
	model "contents-api/internal/models"
	"database/sql"
)

type VodContentRepository struct {
	connection *sql.DB
}

func NewContentVodRepository(connection *sql.DB) *VodContentRepository {
	return &VodContentRepository{
		connection: connection,
	}
}

type Repository interface {
	GetAllVodContentRepository() ([]model.VodContent, error)
	GetVodContentByIDRepository(vodID int) ([]model.VodContent, error)
	CreateVodContentRepository(vodContent model.VodContent) (model.VodContent, error)
}
