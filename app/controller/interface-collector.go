package controller

import "net/http"

type staticController interface {
	HandleStatic(staticDir string) http.Handler
}

type releaseController interface {
	ServeAddTemplate(w http.ResponseWriter, r *http.Request)
	New(w http.ResponseWriter, r *http.Request)
}

type artistController interface {
	New(w http.ResponseWriter, r *http.Request)
}
