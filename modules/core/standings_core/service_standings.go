package standingscore

import (
	core "false_api/modules/core"
	seasoncore "false_api/modules/core/season_core"
	"false_api/modules/models"
	"sync"
)

type StandingsService interface {
	GetStandings(core.Info) (*[]models.Standing, error)
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

func (s standingsSerive) GetStandings(info core.Info) (*[]models.Standing, error) {
	standings, err := s.repo.GetStandings(info.League, info.Season)
	if err != nil {
		return nil, err
	}
	return standings, nil
}

func (s standingsSerive) UpdateStandings(info core.Info) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 5)
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
	for err := range errChan {
		return err
	}
	wg.Done()
	close(errChan)
	close(reserve)

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
