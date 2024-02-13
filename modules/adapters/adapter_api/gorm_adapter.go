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

func (r apiRepository) CreateMatch(home uint, away uint, matchs models.Match) error {
	err := r.db.Where("HomeTeamID = ? AND AwayTeamID = ?", home, away).FirstOrCreate(&matchs)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// _________________________________
// _________________________________
// _________________________________
// _________________________________
// _________________________________
func (r apiRepository) FindLeague(league uint) (uint, error) {
	lea := models.League{}
	err := r.db.Where("LeagueID = ?", league).First(&lea).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return lea.ID, nil
}

func (r apiRepository) FindSeason(season uint) (uint, error) {
	sea := models.Season{}
	err := r.db.Where("Season = ?", season).First(&sea).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return sea.ID, nil
}

func (r apiRepository) CreateLeague(league models.League) (uint, error) {
	if err := r.db.Where(models.League{Name: league.Name}).FirstOrCreate(&league).Error; err != nil {
		return 0, err
	}
	return league.ID, nil
}

func (r apiRepository) CreateSeason(season models.Season) (uint, error) {
	if err := r.db.Where(models.Season{Season: season.Season}).FirstOrCreate(&season).Error; err != nil {
		return 0, err
	}
	return season.Season, nil
}

func (r apiRepository) FindTeam(name string) (uint, error) {
	team := models.Team{}
	err := r.db.Where("Name = ?", name).First(&team).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return team.ID, nil
}

func (r apiRepository) CreateTeam(team models.Team) (uint, error) {
	if err := r.db.Where(models.Team{Name: team.Name}).FirstOrCreate(&team).Error; err != nil {
		return 0, nil
	}
	return team.ID, nil
}

func (r apiRepository) CreatePlayer(player models.Player) error {
	err := r.db.Where("Name = ?", player.Name).FirstOrCreate(&player).Error
	if err != nil {
		return err
	}
	return nil
}
