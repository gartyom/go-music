package service

import (
	"errors"

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

func (as *artist_service) AddNew(artist_name string, image string) (*model.Artist, error) {

	return nil, nil
}

func (as *artist_service) FindByName(artist_name string) (*model.Artist, error) {
	if artist_name == "" {
		return nil, errors.New("artist name is empty")
	}

	_, err := as.repo.Artist.FindName(artist_name)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
