package competitivecore

import (
	"false_api/modules/models"
)

type CompRepository interface {
	GetAllToday() ([]models.Match, error)
}
