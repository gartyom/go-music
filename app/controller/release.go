package controller

import (
	"fmt"
	"log"
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

func (rc *release_controller) New(w http.ResponseWriter, r *http.Request) {

	metadata, err := rc.service.ReleaseForm.Deconstruct(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, item := range *metadata {
		fmt.Println(item.Format.Tags.Title, item.Format.Tags.Album)
	}
	// err = helpers.saveReleaseLocally(unzipper)

	// artists, err := rc.service.Artist.GetByNameMany(release.Artist, "")
	// if err != nil {
	// 	log.Println(err.Error())
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// releaseUuid, err := rc.service.Release.New(artists, release.Title, release.Image)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
}
