package db

import (
	"errors"
	"github.com/kauehmoreno/go-api-rest/times"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//	"log"
)

const TeamCollection = "team"
const PlayerCollection = "player"

var ErrDuplicatedTeam = errors.New("Duplicated team")
var ErrDuplicatedPlayer = errors.New("Duplicated player")

type SoccerRepository struct {
	session *mgo.Session
}

func NewSoccerRepository(session *mgo.Session) *SoccerRepository {
	return &SoccerRepository{session}
}

/*
	Metodos  para TIME
*/
func (r *SoccerRepository) Create(team *times.SoccerTeam) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(TeamCollection)

	err := collection.Insert(team)
	if mongoErr, ok := err.(*mgo.LastError); ok {
		if mongoErr.Code == 11000 {
			return ErrDuplicatedTeam
		}
	}
	return err
}

func (r *SoccerRepository) Update(team *times.SoccerTeam) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(TeamCollection)
	return collection.Update(bson.M{"_id": team.Id}, team)
}

func (r *SoccerRepository) Remove(id string) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(TeamCollection)
	return collection.Remove(bson.M{"_id": id})
}

func (r *SoccerRepository) FindAllActive() ([]*times.SoccerTeam, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(TeamCollection)
	query := bson.M{"inative": false}

	documents := make([]*times.SoccerTeam, 0)

	err := collection.Find(query).All(&documents)
	return documents, err
}

func (r *SoccerRepository) FindById(id string) (*times.SoccerTeam, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(TeamCollection)
	query := bson.M{"_id": id}

	team := &times.SoccerTeam{}

	err := collection.Find(query).One(team)
	return team, err
}

/*
	METODOS PARA PLAYERS
*/

func (r *SoccerRepository) CreatePlayer(player *times.Player) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(PlayerCollection)

	err := collection.Insert(player)
	if mongoErr, ok := err.(*mgo.LastError); ok {
		if mongoErr.Code == 11000 {
			return ErrDuplicatedTeam
		}
	}
	return err
}

func (r *SoccerRepository) UpdatePlayer(player *times.Player) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(PlayerCollection)
	return collection.Update(bson.M{"_id": player.Id}, player)
}

func (r *SoccerRepository) RemovePlayer(id string) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(PlayerCollection)
	return collection.Remove(bson.M{"_id": id})
}

func (r *SoccerRepository) FindAllActivePlayers() ([]*times.Player, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(PlayerCollection)
	query := bson.M{"inative": false}

	documents := make([]*times.Player, 0)

	err := collection.Find(query).All(&documents)
	return documents, err
}

func (r *SoccerRepository) FindPlaylerById(id string) (*times.Player, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(PlayerCollection)
	query := bson.M{"_id": id}

	player := &times.Player{}

	err := collection.Find(query).One(player)
	return player, err
}
