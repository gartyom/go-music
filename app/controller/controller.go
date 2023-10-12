package controller

import (
	"net/http"
	"text/template"

	"github.com/gartyom/go-music/config"
	"github.com/gartyom/go-music/service"
)

type Controller struct {
	Static  staticController
	Release releaseController
	Artist  artistController
}

func Init(service *service.Service) *Controller {
	return &Controller{
		NewStaticController(),
		NewReleaseController(service),
		NewArtistController(service),
	}
}

func (c *Controller) RedirectHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (c *Controller) ServeHomeTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		config.Conf.StaticDir+"/template/index.html",
		config.Conf.StaticDir+"/template/home.html",
		config.Conf.StaticDir+"/template/player.html"))
	tmpl.Execute(w, nil)
}
