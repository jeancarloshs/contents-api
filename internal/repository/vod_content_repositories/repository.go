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
	FindAll() ([]model.VodContent, error)
	FindByID(vodID int) (model.VodContent, error)
	Create(vodContent model.VodContent) (model.VodContent, error)
}
