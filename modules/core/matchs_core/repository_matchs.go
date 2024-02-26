package matchscore

import "false_api/modules/models"

type MatchsRepository interface {
	GetAll(map[string]interface{}) ([]models.Match, error)
	GetOne(map[string]string) (*models.Match, error)
	FindTeam(name string) (uint, error)
	FindLeagueSeason(uint, uint) (uint, error)
	UpdateMatch(models.Match, uint, uint, uint) error
}
