package repository

import (
	"database/sql"

	"github.com/gartyom/go-music/model"
)

type artistPgRepository interface {
	New(artist *model.Artist) (sql.Result, error)
	GetByName(artistName string) (*model.Artist, error)
}

type releasePgRepository interface {
	New(relase *model.Release) (sql.Result, error)
	GetById(releaseUuid string) (*model.Release, error)
	Delete(uuid string) (sql.Result, error)
	UpdateImage(uuid string, image string) (sql.Result, error)
}

type releaseArtistPgRepository interface {
	New(releaseArtist *model.ReleaseArtist) (sql.Result, error)
}
