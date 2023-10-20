package service

import (
	"image"
	"net/http"

	"github.com/gartyom/go-music/model"
)

type artistService interface {
	New(artistName string) (*model.Artist, error)
	GetByName(artistName string) (*model.Artist, error)
	GetByNameMany(artistName string, sep string) ([]*model.Artist, error)
}

type releaseService interface {
	New(artists []*model.Artist, releaseTitle string, cover image.Image) (string, error)
}

type releaseFormService interface {
	Deconstruct(r *http.Request) (*[]Meta, error)
}
