package seasoncore

type SeasonApi interface {
	GetLeague(uint, uint) (*League, error)
	GetStandings(uint, uint) (*Standings, error)
	GetTeam(uint, uint, uint) (*Team, error)
	GetPlayer(uint, uint, int) (*Players, error)
	GetFixture(uint, uint, int) (*Match, error)
}
