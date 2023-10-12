package service

import "github.com/gartyom/go-music/repository"

type release_service struct {
	repo *repository.Repository
}

func NewReleaseService(repo *repository.Repository) releaseService {
	return &release_service{
		repo,
	}
}

func (rs *release_service) AddNew(release_title string, release_cover string) {

}
