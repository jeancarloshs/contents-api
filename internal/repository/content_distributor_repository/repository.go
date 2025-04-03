package content_distributor_repository

import "database/sql"

type ContentDistributorRepository struct {
	connection *sql.DB
}

func NewContentDistributorRepository(connection *sql.DB) ContentDistributorRepository {
	return ContentDistributorRepository{
		connection: connection,
	}
}
