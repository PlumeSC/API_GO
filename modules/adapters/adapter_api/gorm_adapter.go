package adapterapi

import (
	"errors"
	"false_api/modules/models"

	"gorm.io/gorm"
)

type apiRepository struct {
	db *gorm.DB
}

func NewApiRepository(db *gorm.DB) *apiRepository {
	return &apiRepository{db: db}
}

func (r apiRepository) CheckLeague(id uint) (bool, *uint, error) {
	league := models.League{}
	err := r.db.Where("LeagueID = ?", id).First(&league).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, &league.ID, nil
	}
	if err != nil {
		return false, &league.ID, err
	}

	return true, &league.ID, nil
}

func (r apiRepository) CreateLeague(league models.League) (*uint, error) {
	if err := r.db.Where(models.League{Name: league.Name}).FirstOrCreate(&league).Error; err != nil {
		return &league.ID, err
	}
	return &league.ID, nil
}

func (r apiRepository) FindOrCreateSeason(season models.Season) (uint, error) {
	if err := r.db.Where(models.Season{Season: season.Season}).FirstOrCreate(&season).Error; err != nil {
		return 0, nil
	}
	return season.Season, nil
}

func (r apiRepository) CreateTables(standing models.Standing) error {
	result := r.db.Where("TeamID = ? AND SeasonID = ?", standing.TeamID, standing.SeasonID).Attrs(&standing).FirstOrCreate(&standing)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r apiRepository) FindAndCreateTeamName(name string) (uint, error) {
	team := models.Team{}
	if err := r.db.FirstOrCreate(&team, models.Team{Name: name}).Error; err != nil {
		return 0, nil
	}
	return team.ID, nil
}

func (r apiRepository) FindTeam(name string) (*uint, error) {
	team := models.Team{}
	err := r.db.Where("Name = ?", name).First(&team).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &team.ID, nil
	}
	if err != nil {
		return &team.ID, err
	}
	return &team.ID, nil
}

func (r apiRepository) CreateTeam(team models.Team) (*uint, error) {
	if err := r.db.Create(&team).Error; err != nil {
		return &team.ID, nil
	}
	return &team.ID, nil
}

func (r apiRepository) CheckSeason(id uint) (bool, *uint, error) {
	season := models.Season{}
	err := r.db.Where("Season = ?", id).First(&season).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, &season.Season, nil
	}
	if err != nil {
		return false, &season.Season, err
	}
	return true, &season.Season, nil
}
