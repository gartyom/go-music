package controller

import (
	"net/http"

	"github.com/gartyom/go-music/service"
)

type artist_controller struct {
	service *service.Service
}

func NewArtistController(srv *service.Service) artistController {
	return &artist_controller{
		srv,
	}
}

func (ac *artist_controller) New(w http.ResponseWriter, r *http.Request) {
	// ac.service.Artist.AddNew()
}
