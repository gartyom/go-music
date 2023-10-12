package urls

import (
	"github.com/gartyom/go-music/controller"
	"goji.io"
	"goji.io/pat"
)

func GetUrls(mux *goji.Mux, cnt *controller.Controller) {
	mux.HandleFunc(pat.Get("/"), cnt.RedirectHome)
	mux.HandleFunc(pat.Get("/home"), cnt.ServeHomeTemplate)
	releaseUrls(mux, cnt)
	staticUrls(mux, cnt)
}
