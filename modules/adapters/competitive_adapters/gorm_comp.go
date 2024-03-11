package competitiveadapters

import (
	"false_api/modules/models"
	"fmt"
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

	today := time.Now().In(time.UTC)
	err := r.db.Preload("LeagueSeason.League").Preload("LeagueSeason.Season").Where("date(match_day) = ?", today.Format("2006-01-02")).Find(&matches).Error
	if err != nil {
		return nil, err
	}
	for i, v := range matches {

		fmt.Println(i+1, " : ", v.MatchDay)
	}
	return matches, nil
}
