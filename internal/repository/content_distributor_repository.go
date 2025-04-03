package repository

import "database/sql"

type ContentDistributorRepository struct {
	connection *sql.DB
}

func GetContentDistributorRepository(connection *sql.DB) ContentDistributorRepository {
	return ContentDistributorRepository{
		connection: connection,
	}
}
