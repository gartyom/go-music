package repository

import (
	"database/sql"

	"github.com/gartyom/go-music/model"
)

type release_pg_repsitory struct {
	db *sql.DB
}

func NewReleasePgRepo(db *sql.DB) releasePgRepository {
	return &release_pg_repsitory{
		db,
	}
}

func (repo *release_pg_repsitory) GetById(releaseUuid string) (*model.Release, error) {
	var r model.Release
	err := repo.db.QueryRow("SELECT uuid, title, image FROM release WHERE uuid = $1", releaseUuid).Scan(&r.Uuid, &r.Title, &r.Image)
	return &r, err
}

func (repo *release_pg_repsitory) New(r *model.Release) (sql.Result, error) {
	res, err := repo.db.Exec("INSERT INTO release (uuid, title, image) VALUES($1, $2, $3)", r.Uuid, r.Title, r.Image)
	return res, err
}

func (repo *release_pg_repsitory) UpdateImage(uuid string, image string) (sql.Result, error) {
	res, err := repo.db.Exec(`
			UPDATE release 
			SET image = $2 
			WHERE uuid = $1`,
		uuid, image)
	return res, err
}

func (repo *release_pg_repsitory) Delete(uuid string) (sql.Result, error) {
	res, err := repo.db.Exec(`
		DELETE FROM release
		WHERE uuid = $1
	`, uuid)
	return res, err
}
