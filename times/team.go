package times

type SoccerTeam struct {
	Nome      string    `json bson:"nome"`
	Id        string    `json bson:"id"`
	Titulos   string    `json bson:"titulos"`
	Historia  string    `json bson:"historia"`
	Escalacao []*Player `json bson:"escalacao_atual"`
}

type Player struct {
	Nome          string      `json bson:"nome"`
	Posicao       string      `json bson:"posicao"`
	Id            string      `json bson:"id"`
	Gols          int         `json bson:"gols"`
	Cartoes       *Cartao     `json bson:"cartoes"`
	TempoContrato string      `json bson:"tempo_contrato"`
	Nacionalidade string      `json bson:"nacionalidade"`
	Idade         int         `json bson:"idade"`
	Selecao       bool        `json bson:"convocado"`
	contrato      int         `json bson:"tempo_contrato"`
	salario       int         `json bson:"salario"`
	Time          *SoccerTeam `json bson:Time`
}

type Cartao struct {
	Amarelo  int `json bson:"amarelo"`
	Vermelho int `json bson:"vermelho"`
}
