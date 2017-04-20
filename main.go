package main

import (
	"github.com/dimfeld/httptreemux"
	"github.com/kauehmoreno/go-api-rest/api"
	"github.com/kauehmoreno/go-api-rest/db"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	session, err := mgo.Dial("localhost:27017/Soccer2017")

	if err != nil {
		log.Fatal(err)
	}

	repository := db.NewSoccerRepository(session)

	address := "127.0.0.1:8081"
	router := httptreemux.NewContextMux()
	router.Handler(http.MethodGet, "/times/nome/:name", &api.TeamHandler{})
	router.Handler(http.MethodGet, "/time/:name/jogador/:name", &api.PlayerHandler{})

	// posts
	router.Handler(http.MethodPost, "/times/nome/:name/create", &api.PostTeamHandler{Repository: repository})
	router.Handler(http.MethodPost, "/time/:name/jogador/:name/create", &api.PostPlayerHandler{Repository: repository})

	log.Printf("Running api restfull on : http://%s\n", address)
	log.Fatal(http.ListenAndServe(address, router))
}
