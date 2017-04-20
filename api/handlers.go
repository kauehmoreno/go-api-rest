package api

import (
	"encoding/json"
	"fmt"
	"github.com/kauehmoreno/go-api-rest/db"
	"github.com/kauehmoreno/go-api-rest/times"
	"log"
	"net/http"
)

type TeamHandler struct{}

func (team *TeamHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TeamHandler Writer")
}

type PlayerHandler struct{}

func (player *PlayerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PlayerHandler Writer")
}

type PostTeamHandler struct {
	Repository *db.SoccerRepository
}

func (team *PostTeamHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//pegar no contexto
	t := &times.SoccerTeam{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(t)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	error := team.Repository.Create(t)

	if error == db.ErrDuplicatedTeam {
		log.Printf("%s is already Created BITCH\n", t.Nome)
	} else if err != nil {
		log.Printf("Fail to create team:", err)
	}
	fmt.Fprintln(w, "OK")
}

type PostPlayerHandler struct {
	Repository *db.SoccerRepository
}

func (player *PostPlayerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := &times.Player{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	error := player.Repository.CreatePlayer(p)

	if error == db.ErrDuplicatedTeam {
		log.Printf("%s is already Created BITCH\n", p.Nome)
	} else if err != nil {
		log.Printf("Fail to create team:", err)
	}
	fmt.Fprintln(w, "OK")
	//CALMA
}

type AllTeams struct{}

func (allTeams *AllTeams) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "AllTeams writer")
}
