package coreapi

type ApiRequest interface {
	GetTeam(uint, uint, uint) (*Team, error)
	GetLeague(league uint, season uint) (*League, error)
	GetStandings(League uint, season uint) (*Standings, error)
}
