package standingscore

import "false_api/modules/models"

type StandingsRepository interface {
	GetStandings(uint, uint) (*[]models.Standing, error)
	FindTeam(name string) (uint, error)
	FindLeagueSeason(league uint, season uint) (uint, error)
	// UpdateStandings(models.Standing) error
	UpdateStandings(models.Standing, string, uint) error
}
