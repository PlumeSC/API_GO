package matchscore

import (
	seasoncore "false_api/modules/core/season_core"
	"false_api/modules/models"
	"fmt"
	"strconv"
)

type MatchsService interface {
	GetMatchs(params map[string]interface{}) ([]map[string]interface{}, error)
	UpdateMatchs(params map[string]int) error
	GetPlayer(name string) (*models.PlayerStatistics, error)
}

type matchsService struct {
	repo MatchsRepository
	api  seasoncore.SeasonApi
}

func NewMatchService(repo MatchsRepository, api seasoncore.SeasonApi) *matchsService {
	return &matchsService{
		repo: repo,
		api:  api,
	}
}

func (s matchsService) GetMatchs(params map[string]interface{}) ([]map[string]interface{}, error) {
	matches, err := s.repo.GetAll(params)
	if err != nil {
		return nil, err
	}
	data := make([]map[string]interface{}, 0)
	for _, v := range matches {
		data = append(data, map[string]interface{}{
			"HomeTeam":     v.HomeTeam.Name,
			"HomeTeamCode": v.HomeTeam.CodeName,
			"AwayTeam":     v.AwayTeam.Name,
			"AwayTeamCode": v.AwayTeam.CodeName,
			"HomeGoal":     v.HomeGoal,
			"AwayGoal":     v.AwayGoal,
			"Rounded":      v.Rounded,
			"MatchDay":     v.MatchDay,
			"MatchTime":    v.MatchTime,
			"Fixture":      v.Fixture,
			"Season":       v.LeagueSeason.Season.Season,
			"League":       v.LeagueSeason.League.Name,
			"LeagueCode":   v.LeagueSeason.League.ApiCode,
		})
	}
	return data, nil
}

func (s matchsService) UpdateMatchs(params map[string]int) error {
	api, season, round := params["api_code"], params["season"], params["round"]
	match, err := s.api.GetFixture(uint(api), uint(season), round)
	if err != nil {
		return err
	}

	for _, v := range match.Response {

		matchTime := strconv.Itoa(v.Fixture.Status.Elapsed)
		homeID, err := s.repo.FindTeam(v.Teams.Home.Name)
		if err != nil {
			fmt.Println(err)
			return err
		}
		awayID, err := s.repo.FindTeam(v.Teams.Away.Name)
		if err != nil {
			fmt.Println(err)
			return err
		}
		lsID, err := s.repo.FindLeagueSeason(uint(api), uint(season))
		if err != nil {
			fmt.Println(err)
			return err
		}
		data := models.Match{
			Fixture:        uint(v.Fixture.ID),
			HomeGoal:       uint(v.Goals.Home),
			AwayGoal:       uint(v.Goals.Away),
			Rounded:        uint(round),
			MatchDay:       v.Fixture.Date,
			MatchTime:      matchTime,
			HomeTeamID:     homeID,
			AwayTeamID:     awayID,
			LeagueSeasonID: lsID,
		}
		err = s.repo.UpdateMatch(data, homeID, awayID, lsID)
		if err != nil {
			return err
		}
	}
	return err
}

func (s matchsService) GetPlayer(name string) (*models.PlayerStatistics, error) {
	player, err := s.repo.GetPlayer(name)
	if err != nil {
		return nil, err
	}
	return player, nil
}
