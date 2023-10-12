package main

import (
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
	prod := true
	conf := config.Get(prod)

	db := database.Connect(conf)
	defer db.Close()

	repo := repository.Init(db)
	service := service.Init(repo)
	cnt := controller.Init(service)
	mux := goji.NewMux()

	urls.GetUrls(mux, cnt)

	// address := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	// fmt.Println("Listening on address:", address)
	http.ListenAndServe("0.0.0.0:8080", mux)

	return nil
}
