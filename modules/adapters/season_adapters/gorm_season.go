package seasonadapters

import (
	"errors"
	"false_api/modules/models"

	"gorm.io/gorm"
)

type seasonRepository struct {
	db *gorm.DB
}

func NewSeasonRepository(db *gorm.DB) *seasonRepository {
	return &seasonRepository{db: db}
}

func (r seasonRepository) GetLeague(apiCode uint) (uint, error) {
	leagueInfo := models.League{}
	err := r.db.Where("api_code = ?", apiCode).First(&leagueInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return leagueInfo.ID, nil
}

func (r seasonRepository) GetSeason(season uint) (uint, error) {
	seasonInfo := models.Season{}
	err := r.db.Where("season = ?", season).First(&seasonInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return seasonInfo.ID, nil
}

func (r seasonRepository) CreateLeague(leagueInfo models.League) (uint, error) {
	if err := r.db.Create(&leagueInfo).Error; err != nil {
		return 0, nil
	}
	return leagueInfo.ID, nil
}

func (r seasonRepository) CreateSeason(seasonInfo models.Season) (uint, error) {
	if err := r.db.Create(seasonInfo).Error; err != nil {
		return 0, err
	}
	return seasonInfo.ID, nil
}

func (r seasonRepository) LeagueSeasonFirstOrCreate(leagueSeason models.LeagueSeason) (uint, error) {
	err := r.db.Where("season_id = ? AND league_id = ?", leagueSeason.SeasonID, leagueSeason.LeagueID).FirstOrCreate(&leagueSeason).Error
	if err != nil {
		return 0, err
	}
	return leagueSeason.ID, nil
}

func (r seasonRepository) FindTeam(name string) (uint, error) {
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

func (r seasonRepository) CreateTeam(teamInfo models.Team) (uint, error) {
	if err := r.db.Create(&teamInfo).Error; err != nil {
		return 0, err
	}
	return teamInfo.ID, nil
}

func (r seasonRepository) CreateStanding(standingInfo models.Standing) error {
	if err := r.db.Create(&standingInfo).Error; err != nil {
		return err
	}
	return nil
}

func (r seasonRepository) FindOrCreatePlayer(player models.Player) (uint, error) {
	if err := r.db.Where("name = ?", player.Name).FirstOrCreate(&player).Error; err != nil {
		return 0, err
	}
	return player.ID, nil
}

func (r seasonRepository) CreateStatistic(static models.PlayerStatistics) error {
	if err := r.db.Where("player_id = ?", static.PlayerID).Error; err != nil {
		return err
	}
	return nil
}

func (r seasonRepository) CountStandings(leagueseasonID uint) (uint, error) {
	var count int64
	err := r.db.Model(&models.Standing{}).Where("league_season_id = ?", leagueseasonID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return uint(count), nil
}

func (r seasonRepository) CreateMatch(match models.Match) error {
	if err := r.db.Where("home_team_id = ? AND away_team_id = ? AND league_season_id = ?", match.HomeTeamID, match.AwayTeamID, match.LeagueSeasonID).FirstOrCreate(&match).Error; err != nil {
		return err
	}
	return nil
}
