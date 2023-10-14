package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gartyom/go-music/config"
	"github.com/gartyom/go-music/controller"
	"github.com/gartyom/go-music/database"
	"github.com/gartyom/go-music/repository"
	"github.com/gartyom/go-music/service"
	"github.com/gartyom/go-music/urls"
	"goji.io"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	conf := config.Get()

	db := database.Connect(conf)
	defer db.Close()

	repo := repository.Init(db)
	service := service.Init(repo)
	cnt := controller.Init(service)
	mux := goji.NewMux()

	urls.GetUrls(mux, cnt)

	addr := conf.AppHost + ":" + conf.AppPort
	fmt.Println("Listening on address:", addr)
	http.ListenAndServe(addr, mux)

	return nil
}
