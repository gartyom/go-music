package repository

import (
	"database/sql"
	"fmt"

	"github.com/gartyom/go-music/model"
)

type artist_pg_repo struct {
	db *sql.DB
}

func NewArtistSqlRepo(db *sql.DB) artistPgRepository {
	return &artist_pg_repo{
		db,
	}
}

func (repo *artist_pg_repo) Save(artist *model.Artist) (*model.Artist, error) {
	fmt.Println("repository/artist/postgres.go")
	return nil, nil
}

func (repo *artist_pg_repo) FindName(artist_name string) (*model.Artist, error) {
	var artist model.Artist
	err := repo.db.QueryRow("SELECT * FROM artist WHERE name = $1", artist_name).Scan(&artist.Uuid, &artist.Uuid, &artist.Image)
	if err != nil {
		return nil, err
	}
	fmt.Println(artist)
	return nil, nil
}
