package repository

import (
	"database/sql"
)

type Repository struct {
	Artist artistPgRepository
}

func Init(db *sql.DB) *Repository {
	return &Repository{
		NewArtistSqlRepo(db),
	}
}
