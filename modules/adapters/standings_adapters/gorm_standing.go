package standingsadapters

import (
	"errors"
	"false_api/modules/models"

	"gorm.io/gorm"
)

type standingsRepository struct {
	db *gorm.DB
}

func NewstandingsRepository(db *gorm.DB) *standingsRepository {
	return &standingsRepository{db: db}
}

func (r standingsRepository) GetStandings(apiCode uint, season uint) (*[]models.Standing, error) {
	var standing []models.Standing
	var leagueSeason models.LeagueSeason
	if err := r.db.Preload("Season").Preload("League").
		First(&leagueSeason); err.Error != nil {
		return nil, err.Error
	}
	result := r.db.Preload("Team").Preload("LeagueSeason").
		Preload("LeagueSeason.Season").Preload("LeagueSeason.League").
		Where("league_season_id=?", leagueSeason.ID).
		Order("rank").Find(&standing)

	if result.Error != nil {
		return nil, result.Error
	}

	return &standing, nil
}

// func (r standingsRepository) GetStandings(apiCode uint, season uint) (*[]models.Standing, error) {
// 	var standings []models.Standing
// 	result := r.db.Joins("JOIN league_seasons ON league_seasons.league_id = leagues.id").
// 		Joins("JOIN seasons ON league_seasons.season_id = seasons.id").
// 		Joins("JOIN standings ON standings.league_season_id = league_seasons.id").
// 		Where("leagues.api_code = ? AND league_seasons.season = ?", apiCode, season).
// 		Find(&standings)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return nil, errors.New("league or season NotFound")
// 		} else {
// 			return nil, result.Error
// 		}
// 	}
// 	return &standings, nil
// }

func (r standingsRepository) FindTeam(name string) (uint, error) {
	team := models.Team{}
	err := r.db.Where("name = ?", name).First(&team).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return team.ID, nil
}

func (r standingsRepository) FindLeagueSeason(league uint, season uint) (uint, error) {
	var leagueSeason models.LeagueSeason
	result := r.db.Joins("JOIN leagues ON leagues.id = league_seasons.league_id").
		Joins("JOIN seasons ON seasons.id = league_seasons.season_id").
		Where("seasons.season = ? AND leagues.api_code = ?", season, league).
		Select("league_seasons.id").
		First(&leagueSeason).Error
	if result != nil {
		return 0, result
	}
	return leagueSeason.ID, nil
}

func (r standingsRepository) UpdateStandings(standing models.Standing, teamID uint, leagueSeasonID uint) error {

	result := r.db.Model(&models.Standing{}).
		Where("team_id = ? AND league_season_id = ?", teamID, leagueSeasonID).
		Updates(&standing)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
