package matchsadapters

import (
	"false_api/modules/models"

	"gorm.io/gorm"
)

type matchsRepository struct {
	db *gorm.DB
}

func ProvideMatchsRepository(db *gorm.DB) *matchsRepository {
	return &matchsRepository{db: db}
}

func (r *matchsRepository) GetAll(params map[string]string) (*[]models.Match, error) {
	var matches []models.Match
	query := r.db.Model(&models.Match{}).
		Joins("JOIN league_season ON Leagues_season.id = matches.league_season_id").
		// match join league_season ON league_season.id=matches.league_season_id
		Joins("JOIN league ON league.id = league_season.league_id").
		Joins("JOIN season ON season.id = league_season.season_id").
		Joins("LEFT JOIN team AS home_team ON home_team.id = matches.home_team_id").
		Joins("LEFT JOIN team AS away_team ON away_team.id = matches.away_team_id")

	if apiCode, ok := params["api_code"]; ok {
		query = query.Where("league.api_code = ?", apiCode)
	}

	if season, ok := params["season"]; ok {
		query = query.Where("season.season = ?", season)
	}

	if teamName, ok := params["team_name"]; ok {
		query = query.Where("(home_team.name = ? OR away_team.naem = ?)", teamName, teamName)
	}

	if round, ok := params["round"]; ok {
		query = query.Where("match.rounded = ?", round)
	}

	err := query.Find(matches).Error
	if err != nil {
		return nil, err
	}

	return &matches, nil
}

func (r matchsRepository) GetOne(params map[string]string) (*models.Match, error) {
	/*
		params =>
			apicodeLeague,
			season,
			team_away,
			team_home,
	*/
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

type User struct {
	gorm.Model
	Name      string
	AddressID uint
	Address   Address `gorm:"foreignKey:AddressID"`
}

type Address struct {
	gorm.Model
	Street string
	City   string
}

type Team struct {
	gorm.Model

	// Captain Player `gorm:"foreignKey:TeamID"` // One-to-One
}

type Player struct {
	gorm.Model

	TeamID  uint // no foreign key constraint needed
	Captain Team `gorm:"foreignKey:TeamID"` // One-to-One
}
