package controller

import (
	"archive/zip"
	"log"
	"net/http"
	"text/template"

	"github.com/gartyom/go-music/config"
	"github.com/gartyom/go-music/helpers"
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

	release := helpers.ParseReleaseForm(r)
	if release.Err != nil {
		log.Println(release.Err.Error())
		http.Error(w, release.Err.Error(), http.StatusBadRequest)
		return
	}

	fileSize, err := helpers.GetArchiveSize(release.Archive)

	unzipper, err := zip.NewReader(release.Archive, fileSize)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = helpers.CheckArchiveFiles(unzipper)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = helpers.ExtractID3Metadata(unzipper)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
