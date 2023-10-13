package urls

import (
	"github.com/gartyom/go-music/controller"
	"goji.io"
	"goji.io/pat"
)

func releaseUrls(mux *goji.Mux, cnt *controller.Controller) {
	mux.HandleFunc(pat.Post("/release"), cnt.Release.New)
	mux.HandleFunc(pat.Get("/release/add"), cnt.Release.ServeAddTemplate)
}
