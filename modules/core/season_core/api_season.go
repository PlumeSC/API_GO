package seasoncore

import core "false_api/modules/core"

type SeasonApi interface {
	GetLeague(uint, uint) (*core.League, error)
	GetStandings(uint, uint) (*core.Standings, error)
	GetTeam(uint, uint, uint) (*core.Team, error)
	GetPlayer(uint, uint, int) (*core.Players, error)
	GetFixture(uint, uint, int) (*core.Match, error)
}
