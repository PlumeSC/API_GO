package seasoncore

import "false_api/modules/models"

type SeasonRepository interface {
	GetLeague(uint) (uint, error)
	GetSeason(uint) (uint, error)
	CreateLeague(models.League) (uint, error)
	CreateSeason(models.Season) (uint, error)
	LeagueSeasonFirstOrCreate(models.LeagueSeason) (uint, error)
	FindTeam(string) (uint, error)
	CreateTeam(models.Team) (uint, error)
	CreateStanding(models.Standing) error
	FindOrCreatePlayer(models.Player) (uint, error)
	CreateStatistic(models.PlayerStatistics) error
	CountStandings(uint) (uint, error)
	CreateMatch(models.Match) error
}
