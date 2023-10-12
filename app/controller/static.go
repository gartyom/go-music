package controller

import (
	"net/http"
)

type static_controller struct {
}

func NewStaticController() staticController {
	return &static_controller{}
}

func (sc *static_controller) HandleStatic(static_dir string) http.Handler {
	return http.FileServer(http.Dir(static_dir))
}
