package controller

import "net/http"

type staticController interface {
	HandleStatic(static_dir string) http.Handler
}

type releaseController interface {
	ServeAddTemplate(w http.ResponseWriter, r *http.Request)
	AddNew(w http.ResponseWriter, r *http.Request)
}

type artistController interface {
	AddNew(w http.ResponseWriter, r *http.Request)
}
