package service

import "github.com/gartyom/go-music/repository"

type Service struct {
	Artist  artistService
	Release releaseService
}

func Init(repo *repository.Repository) *Service {
	return &Service{
		NewArtistService(repo),
		NewReleaseService(repo),
	}
}
