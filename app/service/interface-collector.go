package service

import (
	"github.com/gartyom/go-music/model"
)

type artistService interface {
	AddNew(artist_name string, image string) (*model.Artist, error)
	FindByName(artist_name string) (*model.Artist, error)
}

type releaseService interface {
	AddNew(release_title string, release_cover string)
}
