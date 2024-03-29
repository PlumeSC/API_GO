package competitiveadapters

import (
	"false_api/modules/models"
	"time"

	"gorm.io/gorm"
)

type compRepository struct {
	db *gorm.DB
}

func NewCompRepository(db *gorm.DB) *compRepository {
	return &compRepository{db: db}
}

func (r compRepository) GetAllToday() ([]models.Match, error) {
	var matches []models.Match

	// location, _ := time.LoadLocation("UTC")
	// today := time.Date(2024, time.March, 2, 0, 0, 0, 0, location)

	today := time.Now()
	err := r.db.Preload("LeagueSeason.League").Preload("LeagueSeason.Season").Where("date(match_day) = ?", today.Format("2006-01-02")).Find(&matches).Error
	if err != nil {
		return nil, err
	}
	return matches, nil
}
