package repository

import (
	"database/sql"

	"github.com/gartyom/go-music/model"
)

type release_artist_pg_repo struct {
	db *sql.DB
}

func NewReleaseArtistPgRepo(db *sql.DB) releaseArtistPgRepository {
	return &release_artist_pg_repo{
		db,
	}
}

func (repo *release_artist_pg_repo) New(release_artist *model.ReleaseArtist) (sql.Result, error) {
	res, err := repo.db.Exec(`
		INSERT INTO release_artist (artist_uuid, release_uuid) 
		VALUES($1, $2)
	`,
		release_artist.ArtistUuid,
		release_artist.ReleaseUuid)
	return res, err
}
