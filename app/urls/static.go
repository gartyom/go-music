package urls

import (
	"net/http"

	"github.com/gartyom/go-music/config"
	"github.com/gartyom/go-music/controller"
	"goji.io"
	"goji.io/pat"
)

func staticUrls(mux *goji.Mux, cnt *controller.Controller) {
	mux.Handle(pat.Get("/static/*"), http.StripPrefix("/static/", cnt.Static.HandleStatic(config.Conf.StaticDir)))
}
