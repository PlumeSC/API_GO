package seasoncore

import (
	"errors"
	core "false_api/modules/core"
	"false_api/modules/models"
	"strconv"
	"sync"
	"time"
)

type SeasonService interface {
	CreateStandings(core.Info) error
	CreatePlayers(core.Info) error
	CreateMatch(core.Info) error
}

type seasonService struct {
	repo SeasonRepository
	api  SeasonApi
}

func NewSeasonService(repo SeasonRepository, api SeasonApi) *seasonService {
	return &seasonService{
		repo: repo,
		api:  api,
	}
}

func (s seasonService) CreateStandings(info core.Info) error {
	id, leagueID, err := s.getLeagueSeason(info)
	if err != nil {
		return err
	}
	err = s.createStandings(id, info, leagueID)
	if err != nil {
		return err
	}
	return nil
}

func (s seasonService) getLeagueSeason(info core.Info) (uint, uint, error) {
	var wg sync.WaitGroup
	leagueChan := make(chan uint)
	seasonChan := make(chan uint)
	errChan := make(chan error)

	wg.Add(1)
	go s.getLeague(&wg, info, leagueChan, errChan)
	wg.Add(1)
	go s.getSeason(&wg, info, seasonChan, errChan)

	leagueID := <-leagueChan
	seasonID := <-seasonChan
	wg.Wait()
	close(leagueChan)
	close(seasonChan)
	close(errChan)
	leagueSeason := models.LeagueSeason{
		SeasonID: seasonID,
		LeagueID: leagueID,
	}
	leagueSeasonID, err := s.repo.LeagueSeasonFirstOrCreate(leagueSeason)
	if err != nil {
		return 0, 0, err
	}
	return leagueSeasonID, leagueID, nil
}

func (s seasonService) getLeague(wg *sync.WaitGroup, info core.Info, leagueChan chan<- uint, errChan chan<- error) {
	defer wg.Done()
	id, err := s.repo.GetLeague(info.League)
	if err != nil {
		errChan <- err
	}
	if id == 0 {
		id, err = s.createLeague(info)
		if err != nil {
			errChan <- err
		}
	}
	leagueChan <- id
}

func (s seasonService) createLeague(info core.Info) (uint, error) {
	leagueInfo, err := s.api.GetLeague(info.League, info.Season)
	if err != nil {
		return 0, err
	}
	x := leagueInfo.Response[0]
	league := models.League{
		Name:    x.League.Name,
		Country: x.Country.Name,
		Logo:    x.League.Logo,
		Flag:    x.Country.Flag,
		ApiCode: uint(x.League.ID),
	}
	id, err := s.repo.CreateLeague(league)
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (s seasonService) getSeason(wg *sync.WaitGroup, info core.Info, seasonChan chan<- uint, errChan chan<- error) {
	defer wg.Done()
	id, err := s.repo.GetSeason(info.Season)
	if err != nil {
		errChan <- err
	}
	if id == 0 {
		id, err = s.createSeason(info)
		if err != nil {
			errChan <- err
		}
	}
	seasonChan <- id
}

func (s seasonService) createSeason(info core.Info) (uint, error) {
	leagueInfo, err := s.api.GetLeague(info.League, info.Season)
	if err != nil {
		return 0, err
	}

	layout := "2006-01-02"
	start, err := time.Parse(layout, leagueInfo.Response[0].Seasons[0].Start)
	if err != nil {
		return 0, err
	}
	end, err := time.Parse(layout, leagueInfo.Response[0].Seasons[0].End)
	if err != nil {
		return 0, err
	}

	season := models.Season{
		Season:  uint(leagueInfo.Response[0].Seasons[0].Year),
		Start:   start,
		End:     end,
		Current: leagueInfo.Response[0].Seasons[0].Current,
	}
	id, err := s.repo.CreateSeason(season)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s seasonService) createStandings(leagueSeasonID uint, info core.Info, leagueID uint) error {
	var wg sync.WaitGroup
	errChan := make(chan error)
	reserve := make(chan struct{}, 5)

	standingsInfo, err := s.api.GetStandings(info.League, info.Season)
	if err != nil {
		return err
	}

	for i, v := range standingsInfo.Response[0].League.Standings[0] {
		wg.Add(1)
		go func(i int, v core.Standing) {
			defer wg.Done()
			reserve <- struct{}{}
			defer func() { <-reserve }()
			teamID, err := s.firstOrCreateTeam(v.Team.Name, uint(v.Team.ID), info, leagueID)
			if err != nil {
				errChan <- err
			}
			err = s.createStanding(teamID, uint(v.Rank), leagueSeasonID)
			if err != nil {
				errChan <- err
			}
		}(i, v)
	}

	return nil
}

func (s seasonService) firstOrCreateTeam(name string, teamApi uint, info core.Info, leagueID uint) (uint, error) {
	teamID, err := s.repo.FindTeam(name)
	if err != nil {
		return 0, err
	}
	if teamID == 0 {
		teamInfo, err := s.api.GetTeam(teamApi, info.League, info.Season)
		if err != nil {
			return 0, nil
		}
		res := teamInfo.Response[0]
		team := models.Team{
			Name:        res.Team.Name,
			CodeName:    res.Team.Code,
			Founded:     uint(res.Team.Founded),
			Logo:        res.Team.Logo,
			StadiumName: res.Venue.Name,
			City:        res.Venue.City,
			Capacity:    uint(res.Venue.Capacity),
			LeagueID:    leagueID,
		}
		teamID, err = s.repo.CreateTeam(team)
		if err != nil {
			return 0, err
		}
	}
	return teamID, nil
}

func (s seasonService) createStanding(teamID uint, rank uint, leagueSeasonID uint) error {
	standing := models.Standing{
		Rank:           rank,
		Form:           "",
		TeamID:         teamID,
		LeagueSeasonID: leagueSeasonID,
	}
	err := s.repo.CreateStanding(standing)
	if err != nil {
		return err
	}
	return nil
}

func (s seasonService) CreatePlayers(info core.Info) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 5)
	reserve := make(chan struct{}, 4)
	totalPages := 1

	go func() {
		for currentPages := 1; currentPages <= totalPages; currentPages++ {
			wg.Add(1)
			go func(page int) {
				defer wg.Done()
				reserve <- struct{}{}
				defer func() { <-reserve }()

				players, err := s.api.GetPlayer(info.League, info.Season, page)
				if err != nil {
					errChan <- err
				}

				if page == 1 && totalPages == 1 {
					totalPages = players.Paging.Total
				}
				err = s.createPlayerAndPlayerStatistics(*players)
				if err != nil {
					errChan <- err
				}

			}(currentPages)
		}
	}()
	go func() {
		wg.Wait()
		close(reserve)
		close(errChan)
	}()
	for err := range errChan {
		return err
	}
	return nil
}

func (s seasonService) createPlayerAndPlayerStatistics(players core.Players) error {
	for _, v := range players.Response {
		TeamID, err := s.repo.FindTeam(v.Statistics[0].Team.Name)
		if err != nil {
			return err
		}
		height, err := strconv.Atoi(v.Player.Height)
		if err != nil {
			return err
		}
		weight, err := strconv.Atoi(v.Player.Weight)
		if err != nil {
			return err
		}
		player := models.Player{
			Name:        v.Player.Name,
			Firstname:   v.Player.Firstname,
			Lastname:    v.Player.Lastname,
			Age:         uint(v.Player.Age),
			Nationality: v.Player.Nationality,
			Height:      uint(height),
			Weight:      uint(weight),
			Injuries:    v.Player.Injured,
			Photo:       v.Player.Photo,
			TeamID:      TeamID,
		}
		playerID, err := s.repo.FindOrCreatePlayer(player)
		if err != nil {
			return err
		}
		number, ok := v.Statistics[0].Games.Number.(float64)
		if !ok {
			return errors.New("expected type float64")
		}
		static := models.PlayerStatistics{
			PlayerID: playerID,
			Number:   uint(number),
			Position: v.Statistics[0].Games.Position,
		}
		err = s.repo.CreateStatistic(static)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s seasonService) CreateMatch(info core.Info) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 4)
	reserve := make(chan struct{}, 4)

	leagueSeasonID, _, err := s.getLeagueSeason(info)
	if err != nil {
		return err
	}
	round, err := s.repo.CountStandings(leagueSeasonID)
	if err != nil {
		return err
	}
	round2 := (round - 1) * 2
	for i := 1; i <= int(round2); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			reserve <- struct{}{}
			defer func() { <-reserve }()

			matchs, err := s.api.GetFixture(info.League, info.Season, i)
			if err != nil {
				errChan <- err
			}
			for j := 0; j < len(matchs.Response); j++ {
				res := matchs.Response[j]
				homeTeamID, err := s.repo.FindTeam(res.Teams.Home.Name)
				if err != nil {
					errChan <- err
					return
				}
				awayTeamID, err := s.repo.FindTeam(res.Teams.Away.Name)
				if err != nil {
					errChan <- err
					return
				}
				match := models.Match{
					HomeTeamID:     homeTeamID,
					AwayTeamID:     awayTeamID,
					LeagueSeasonID: leagueSeasonID,
					Rounded:        uint(i),
					MatchDay:       res.Fixture.Date,
				}
				err = s.repo.CreateMatch(match)
				if err != nil {
					errChan <- err
				}
			}

		}(i)
	}
	go func() {
		wg.Wait()
		close(errChan)
		close(reserve)
	}()
	for err := range errChan {
		return err
	}
	return nil
}
