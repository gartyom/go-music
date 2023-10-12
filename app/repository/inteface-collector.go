package repository

import (
	"github.com/gartyom/go-music/model"
)

type artistPgRepository interface {
	Save(artist *model.Artist) (*model.Artist, error)
	FindName(artist_name string) (*model.Artist, error)
}
