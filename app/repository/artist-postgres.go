package repository

import (
	"database/sql"
	"fmt"

	"github.com/gartyom/go-music/model"
)

type artist_pg_repo struct {
	db *sql.DB
}

func NewArtistPgRepo(db *sql.DB) artistPgRepository {
	return &artist_pg_repo{
		db,
	}
}

func (repo *artist_pg_repo) New(artist *model.Artist) (sql.Result, error) {
	fmt.Println("repository/artist/postgres.go")
	return nil, nil
}

func (repo *artist_pg_repo) GetByName(artist_name string) (*model.Artist, error) {
	var artist model.Artist
	err := repo.db.QueryRow("SELECT * FROM artist WHERE name = $1", artist_name).Scan(&artist.Uuid, &artist.Name, &artist.Image)
	return &artist, err
}
