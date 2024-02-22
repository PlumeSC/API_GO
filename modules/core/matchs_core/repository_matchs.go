package matchscore

import "false_api/modules/models"

type MatchsRepository interface {
	GetAll(map[string]string) (*[]models.Match, error)
	GetOne(map[string]string) (*models.Match, error)
	UpdateMatch(*models.Match) error
}
