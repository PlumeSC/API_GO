package coreapi

import (
	"false_api/modules/models"
)

type ApiRepository interface {
	FindOrCreateSeason(models.Season) (uint, error)
	CreateTables(models.Standing) error
	FindAndCreateTeamName(string) (uint, error)
	// CheckLeagueSeason(uint, uint) (bool, error)
	CheckLeague(uint) (bool, *uint, error)
	CheckSeason(uint) (bool, *uint, error)

	CreatePlayer(models.Player) error
	CreateMatch(uint, uint, models.Match) error

	FindLeague(uint) (uint, error)
	FindSeason(uint) (uint, error)
	CreateLeague(models.League) (uint, error)
	CreateSeason(models.Season) (uint, error)
	FindTeam(string) (uint, error)
	CreateTeam(models.Team) (uint, error)
}
