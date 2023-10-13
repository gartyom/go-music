package repository

import (
	"database/sql"
)

type Repository struct {
	Artist       artistPgRepository
	Release      releasePgRepository
	RelaseArtist releaseArtistPgRepository
}

func Init(db *sql.DB) *Repository {
	return &Repository{
		NewArtistPgRepo(db),
		NewReleasePgRepo(db),
		NewReleaseArtistPgRepo(db),
	}
}
