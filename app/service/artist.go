package service

import (
	"strings"

	"github.com/gartyom/go-music/model"
	"github.com/gartyom/go-music/repository"
)

type artist_service struct {
	repo *repository.Repository
}

func NewArtistService(repo *repository.Repository) artistService {
	return &artist_service{
		repo,
	}
}

func (as *artist_service) New(artist_name string, image string) (*model.Artist, error) {

	return nil, nil
}

func (as *artist_service) GetByName(artist_name string) (*model.Artist, error) {
	artist, err := as.repo.Artist.GetByName(artist_name)
	return artist, err
}

func (as *artist_service) GetByNameMany(artist_name string, sep string) ([]*model.Artist, error) {

	if sep == "" {
		sep = "/"
	}

	var a []*model.Artist
	for _, v := range strings.Split(artist_name, sep) {
		v = strings.TrimSpace(v)
		artist, err := as.repo.Artist.GetByName(v)
		if err != nil {
			return nil, err
		}
		a = append(a, artist)
	}

	return a, nil
}
