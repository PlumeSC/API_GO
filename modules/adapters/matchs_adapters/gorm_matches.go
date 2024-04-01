package matchsadapters

import (
	"false_api/modules/models"
	"fmt"

	"gorm.io/gorm"
)

type matchsRepository struct {
	db *gorm.DB
}

func NewMatchsRepository(db *gorm.DB) *matchsRepository {
	return &matchsRepository{db: db}
}

func (r matchsRepository) GetAll(params map[string]interface{}) ([]models.Match, error) {
	var matchs []models.Match

	query := r.db.Model(&models.Match{}).
		Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("LeagueSeason.League").
		Preload("LeagueSeason.Season").
		Joins("join teams as t on matches.home_team_id = t.id").
		Joins("join teams as t2 on matches.away_team_id = t2.id").
		Joins("join league_seasons as ls on matches.league_season_id = ls.id").
		Joins("join leagues as l on ls.league_id = l.id").
		Joins("join seasons as s on ls.season_id = s.id")

	if apiCode, ok := params["api_code"]; ok {
		query = query.Where("l.api_code = ?", apiCode)
	}

	if season, ok := params["season"]; ok {
		query = query.Where("s.season = ?", season)
	}
	if teamName, ok := params["team_name"]; ok {
		if teamName != "" {
			query = query.Where("t.name = ? OR t2.name = ?", teamName, teamName)
		}
	}

	if roundValue, ok := params["round"]; ok {
		if round, ok := roundValue.(uint); ok {
			if round != 0 {
				query = query.Where("matches.rounded = ?", round)
			}
		}
	}

	err := query.Order("rounded").Find(&matchs).Error
	if err != nil {
		return nil, err
	}

	return matchs, nil
}

func (r matchsRepository) GetAll2(params map[string]interface{}) ([]models.Match, error) {
	var matches []models.Match

	var ls models.LeagueSeason
	lsE := r.db.Joins("JOIN leagues ON leagues.id = league_seasons.league_id").
		Joins("JOIN seasons ON seasons.id = league_seasons.season_id").
		Where("leagues.api_code = ? AND seasons.season = ?", params["api_code"], params["season"]).
		First(&ls).Error
	if lsE != nil {
		return nil, lsE
	}
	fmt.Println(ls.ID)

	var team models.Team
	teamE := r.db.Where("name = ?", params["team_name"]).First(&team).Error
	if teamE != nil {
		return nil, teamE
	}
	fmt.Println(team.ID)

	query := r.db.Model(&models.Match{}).
		Where("league_season_id = ? AND (home_team_id = ? OR away_team_id = ?)", ls.ID, team.ID, team.ID).
		Find(&matches).Error

	if query != nil {
		return nil, query
	}

	return matches, nil
}

func (r matchsRepository) GetOne(params map[string]string) (*models.Match, error) {
	var matchs models.Match
	query := r.db.Model(&models.Match{}).
		Joins("JOIN league_season ON league_season.id = match.league_season_id").
		Joins("JOIN league ON league.id = league_season.league_id").
		Joins("JOIN season ON season.id = league_season.season_id").
		Joins("JOin ").
		Joins("").
		Joins("LEFT JOIN team AS home_team ON home_team.id = match.home_team_id").
		Joins("LEFT JOIN team AS away_team ON away_team.id = match.away_team_id")

	_ = query
	return &matchs, nil
}

func (r matchsRepository) FindTeam(name string) (uint, error) {
	var team models.Team
	if err := r.db.Where("name = ?", name).First(&team).Error; err != nil {
		return 0, err
	}
	return team.ID, nil
}
func (r matchsRepository) FindLeagueSeason(api uint, season uint) (uint, error) {
	var leagueSeason models.LeagueSeason
	if err := r.db.Joins("join leagues on league_seasons.league_id = leagues.id").
		Joins("join seasons on league_seasons.season_id = seasons.id").
		Where("leagues.api_code = ? AND seasons.season =?", api, season).
		First(&leagueSeason).Error; err != nil {
		return 0, err
	}
	return leagueSeason.ID, nil
}

func (r matchsRepository) UpdateMatch(match models.Match, homeID uint, awayID uint, lsID uint) error {
	if err := r.db.Where("home_team_id = ? AND away_team_id = ? AND league_season_id = ?", homeID, awayID, lsID).
		Updates(&match).Error; err != nil {
		return err
	}
	return nil
}

func (r matchsRepository) GetPlayer(name string) (*models.PlayerStatistics, error) {
	player := models.PlayerStatistics{}
	err := r.db.Preload("Player").Preload("Player.Team").Joins("FULL OUTER JOIN players on player_statistics.player_id = players.id").Where("players.name = ?", name).First(&player).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(player)
	return &player, nil
}
