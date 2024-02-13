package coreapi

type ApiRequest interface {
	GetTeam(uint, uint, uint) (*Team, error)
	GetLeague(league uint, season uint) (*League, error)
	GetStandings(League uint, season uint) (*Standings, error)
	GetPlayer(League uint, season uint, page int) (*Player, error)
	GetFixture(League uint, season uint, round int) (*Match, error)
}
