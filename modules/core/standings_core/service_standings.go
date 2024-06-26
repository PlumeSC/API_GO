package standingscore

import (
	core "false_api/modules/core"
	seasoncore "false_api/modules/core/season_core"
	"false_api/modules/models"
	"fmt"
	"sync"
)

type StandingsService interface {
	GetStandings(core.Info) ([]map[string]interface{}, error)
	UpdateStandings(core.Info) error
}

type standingsSerive struct {
	repo StandingsRepository
	Api  seasoncore.SeasonApi
}

func NewStandingsService(repo StandingsRepository, Api seasoncore.SeasonApi) *standingsSerive {
	return &standingsSerive{
		repo: repo,
		Api:  Api,
	}
}

func (s standingsSerive) GetStandings(info core.Info) ([]map[string]interface{}, error) {
	standings, err := s.repo.GetStandings(info.League, info.Season)
	if err != nil {
		return nil, err
	}
	data := make([]map[string]interface{}, 0)
	for _, v := range *standings {
		data = append(data, map[string]interface{}{
			"Team":     v.Team.Name,
			"CodeName": v.Team.CodeName,
			"TeamImg":  v.Team.Logo,
			"Rank":     v.Rank,
			"Played":   v.Played,
			"win":      v.Won,
			"Draw":     v.Drawn,
			"Lost":     v.Lost,
			"GF":       v.GF,
			"GA":       v.GA,
			"GD":       v.GD,
			"Points":   v.Points,
			"Form":     v.Form,
		})
	}
	return data, nil
}

func (s standingsSerive) UpdateStandings(info core.Info) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 4)
	reserve := make(chan struct{}, 2)

	leagueSeasonID, err := s.repo.FindLeagueSeason(info.League, info.Season)
	if err != nil {
		return err
	}

	standingsApi, err := s.Api.GetStandings(info.League, info.Season)
	if err != nil {
		return err
	}
	for i, v := range standingsApi.Response[0].League.Standings[0] {
		wg.Add(1)
		go func(i int, v core.Standing) {
			defer wg.Done()
			reserve <- struct{}{}
			defer func() { <-reserve }()
			err = s.updateStandings(v, v.Team.Name, leagueSeasonID)
			if err != nil {
				errChan <- err
			}
		}(i, v)
	}

	fmt.Println("ok")
	wg.Wait()
	close(errChan)
	close(reserve)
	for err := range errChan {
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (s standingsSerive) updateStandings(standing core.Standing, teamName string, leagueSeasonID uint) error {
	teamID, err := s.repo.FindTeam(teamName)
	if err != nil {
		return err
	}
	updatestanding := models.Standing{
		Rank:           uint(standing.Rank),
		Played:         uint(standing.All.Played),
		Won:            uint(standing.All.Win),
		Drawn:          uint(standing.All.Draw),
		Lost:           uint(standing.All.Lose),
		GF:             uint(standing.All.Goals.For),
		GA:             uint(standing.All.Goals.Against),
		GD:             standing.GoalsDiff,
		Points:         uint(standing.Points),
		Form:           standing.Form,
		LeagueSeasonID: leagueSeasonID,
		TeamID:         teamID,
	}

	err = s.repo.UpdateStandings(updatestanding, teamID, leagueSeasonID)
	if err != nil {
		return err
	}

	return nil
}
