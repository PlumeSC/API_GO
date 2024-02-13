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

func (r apiRepository) CreateTables(standing models.Standing) error {
	result := r.db.Where("TeamID = ? AND SeasonID = ?", standing.TeamID, standing.SeasonID).Attrs(&standing).FirstOrCreate(&standing)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r apiRepository) CreatePlayer(player models.Player) (uint, error) {
	err := r.db.Where("Name = ?", player.Name).FirstOrCreate(&player).Error
	if err != nil {
		return 0, err
	}
	return player.ID, nil
}

func (r apiRepository) CreatePlayerStatistics(id uint, statistics models.PlayerStatistics) error {
	err := r.db.Where("PlayerID = ?", id).FirstOrCreate(&statistics).Error
	if err != nil {
		return err
	}
	return nil
}

func (r apiRepository) CountStandings(league uint, season uint) (uint, error) {
	var count int64
	leagueID, err := r.FindLeague(league)
	if err != nil {
		return 0, err
	}
	seasonID, err := r.FindSeasonByLeague(leagueID, season)
	if err != nil {
		return 0, err
	}
	err = r.db.Model(&models.Standing{}).Where("seasonID = ?", seasonID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return uint(count), nil
}

func (r apiRepository) FindSeasonByLeague(leagueID uint, season uint) (uint, error) {
	seasonX := models.Season{}
	err := r.db.Where("season = ? AND LeagueID = ?", season, leagueID).First(&seasonX).Error
	if err != nil {
		return 0, err
	}
	return seasonX.ID, nil
}

func (r apiRepository) CreateMatch(matchs models.Match) error {
	err := r.db.Where("HomeTeamID = ? AND AwayTeamID = ? AND seasonID = ?", matchs.HomeTeamID, matchs.AwayTeamID, matchs.SeasonID).FirstOrCreate(&matchs)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
