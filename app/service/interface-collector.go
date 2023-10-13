package service

import (
	"image"

	"github.com/gartyom/go-music/model"
)

type artistService interface {
	New(artist_name string, image string) (*model.Artist, error)
	GetByName(artist_name string) (*model.Artist, error)
	GetByNameMany(artist_name string, sep string) ([]*model.Artist, error)
}

type releaseService interface {
	New(artists []*model.Artist, release_title string, cover image.Image) error
}
