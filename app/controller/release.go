package controller

import (
	"archive/zip"
	"image/jpeg"
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

	var err error
	artist_name := r.FormValue("artist")
	if artist_name == "" {
		s := "Error: artist name is missing"
		log.Println(s)
		http.Error(w, s, http.StatusBadRequest)
		return
	}

	release_title := r.FormValue("release")
	if release_title == "" {
		s := "Error: release title is missing"
		log.Println(s)
		http.Error(w, s, http.StatusBadRequest)
		return
	}

	cover, _, err := r.FormFile("release_cover")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	archive, _, err := r.FormFile("release_songs")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	image, err := jpeg.Decode(cover)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Finding size of archive
	fileSize, err := archive.Seek(0, 2)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = archive.Seek(0, 0)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	unzipper, err := zip.NewReader(archive, fileSize)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// read files from unzipper
	// check if mp3
	// if ok send to service.Song.NewMany or something

	artists, err := rc.service.Artist.GetByNameMany(artist_name, "")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = rc.service.Release.New(artists, release_title, image)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
