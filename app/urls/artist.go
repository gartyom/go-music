package urls

import (
	"github.com/gartyom/go-music/controller"
	"goji.io"
	"goji.io/pat"
)

func artistUrls(mux *goji.Mux, cnt *controller.Controller) {
	mux.HandleFunc(pat.Post("/artist"), cnt.Artist.New)
}
