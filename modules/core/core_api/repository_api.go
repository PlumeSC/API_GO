package coreapi

import (
	"false_api/modules/models"
)

type ApiRepository interface {
	FindLeague(uint) (uint, error)
	FindSeason(uint) (uint, error)
	CreateLeague(models.League) (uint, error)
	CreateSeason(models.Season) (uint, error)
	FindTeam(string) (uint, error)
	CreateTeam(models.Team) (uint, error)
	CreateTables(models.Standing) error

	CreatePlayer(models.Player) (uint, error)
	CreatePlayerStatistics(uint, models.PlayerStatistics) error

	CountStandings(uint, uint) (uint, error)
	CreateMatch(models.Match) error
}
