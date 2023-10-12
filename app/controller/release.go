package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gartyom/go-music/config"
	"github.com/gartyom/go-music/service"
)

type release_controller struct {
	service *service.Service
}

func NewReleaseController(srv *service.Service) releaseController {
	return &release_controller{
		srv,
	}
}

func (rc *release_controller) ServeAddTemplate(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	if r.Header.Get("HX-Request") == "true" {
		tmpl = template.Must(template.ParseFiles(config.Conf.StaticDir + "/template/release/add.html"))
	} else {
		tmpl = template.Must(template.ParseFiles(
			config.Conf.StaticDir+"/template/index.html",
			config.Conf.StaticDir+"/template/player.html",
			config.Conf.StaticDir+"/template/release/add.html",
		))
	}
	tmpl.Execute(w, nil)
}

func (rc *release_controller) AddNew(w http.ResponseWriter, r *http.Request) {
	// release := r.FormValue("release")
	artist_name := r.FormValue("artist")

	fmt.Println(artist_name)

	_, err := rc.service.Artist.FindByName(artist_name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// cover, cheader, err := r.FormFile("release_cover")
	// defer cover.Close()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// rc.service.Release.AddNew(release)
}
