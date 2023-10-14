package controller

import (
	"net/http"
)

type static_controller struct {
}

func NewStaticController() staticController {
	return &static_controller{}
}

func (sc *static_controller) HandleStatic(staticDir string) http.Handler {
	return http.FileServer(http.Dir(staticDir))
}
