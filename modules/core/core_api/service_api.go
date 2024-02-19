package coreapi

import (
	"errors"
	"false_api/modules/models"
	"strconv"
	"sync"
	"time"
)

type ApiService interface {
	CreateStandings(Info) error
	CreatePlayers(Info) error
	CreateMatch(Info) error
}

type apiService struct {
	repo ApiRepository
	api  ApiRequest
}

func NewApiService(repo ApiRepository, api ApiRequest) *apiService {
	return &apiService{
		repo: repo,
		api:  api,
	}
}

func (s apiService) goFirstOrCreateLeague(info Info, wg *sync.WaitGroup, leagueChan chan<- uint, leagueStructChan <-chan League, errChan chan<- error) {
	defer wg.Done()

	id, err := s.repo.FindLeague(info.League)
	if err != nil {
		errChan <- err
		return
	}
	if id == 0 {
		league := <-leagueStructChan
		id, err := s.CreateLeague(league)
		if err != nil {
			errChan <- err
			return
		}
		leagueChan <- id
		leagueChan <- id
		return
	}
	leagueChan <- id
}
func (s apiService) CreateLeague(league League) (uint, error) {
	x := league.Response[0]
	leagueStruct := models.League{
		Name:     x.League.Name,
		LeagueID: uint(x.League.ID),
		Country:  x.Country.Name,
		Code:     x.Country.Code,
		Logo:     x.League.Logo,
		Flag:     x.Country.Flag,
	}
	leagueID, err := s.repo.CreateLeague(leagueStruct)
	if err != nil {
		return 0, err
	}
	return leagueID, nil
}
func (s apiService) goFirstOrCreateSeason(info Info, wg *sync.WaitGroup, seasonChan chan<- uint, leagueChan <-chan uint, leagueStructChan chan<- League, errChan chan<- error) {
	defer wg.Done()

	id, err := s.repo.FindSeason(info.Season)
	if err != nil {
		errChan <- err
		return
	}

	if id == 0 {
		league := &League{}
		leagueID, err := s.repo.FindLeague(info.League)
		if err != nil {
			errChan <- err
			return
		}
		if leagueID == 0 {
			league, err = s.api.GetLeague(info.League, info.Season)
			if err != nil {
				errChan <- err
				return
			}
			leagueStructChan <- *league
			leagueID = <-leagueChan

		}
		id, err := s.CreateSeason(*league, &leagueID)
		if err != nil {
			errChan <- err
			return
		}
		seasonChan <- id
		return
	}
	seasonChan <- id
}
func (s apiService) CreateSeason(league League, leagueID *uint) (uint, error) {
	layout := "2006-01-02"
	start, err := time.Parse(layout, league.Response[0].Seasons[0].Start)
	if err != nil {
		return 0, err
	}
	end, err := time.Parse(layout, league.Response[0].Seasons[0].End)
	if err != nil {
		return 0, err
	}
	season := models.Season{
		Season:   uint(league.Response[0].Seasons[0].Year),
		LeagueID: *leagueID,
		Start:    start,
		End:      end,
		Current:  league.Response[0].Seasons[0].Current,
	}
	seasonID, err := s.repo.CreateSeason(season)
	if err != nil {
		return 0, err
	}
	return seasonID, nil
}
func (s apiService) getLeagueAndSeason(info Info) (uint, uint, error) {
	var wg sync.WaitGroup
	leagueChan := make(chan uint)
	seasonChan := make(chan uint)
	leagueStructChan := make(chan League)
	errChan := make(chan error)

	//league
	wg.Add(1)
	go s.goFirstOrCreateLeague(info, &wg, leagueChan, leagueStructChan, errChan)
	// season
	wg.Add(1)
	go s.goFirstOrCreateSeason(info, &wg, seasonChan, leagueChan, leagueStructChan, errChan)

	wg.Wait()
	leagueID := <-leagueChan
	seasonID := <-seasonChan
	close(leagueChan)
	close(seasonChan)
	close(leagueStructChan)
	close(errChan)

	for err := range errChan {
		return 0, 0, err
	}

	return leagueID, seasonID, nil
}
func (s apiService) goFirstOrCreateTeam(name string, teamID_Api uint, leagueID uint, seasonID uint) (uint, error) {
	id, err := s.repo.FindTeam(name)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		team, err := s.api.GetTeam(teamID_Api, leagueID, seasonID)
		if err != nil {
			return 0, nil
		}
		id, err = s.createTeam(*team, leagueID)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}
func (s apiService) createTeam(team Team, leagueID uint) (uint, error) {
	res := team.Response[0]
	teamInfo := models.Team{
		Name:        res.Team.Name,
		CodeName:    res.Team.Code,
		Founded:     uint(res.Team.Founded),
		Logo:        res.Team.Logo,
		StadiumName: res.Venue.Name,
		City:        res.Venue.City,
		Capacity:    uint(res.Venue.Capacity),
		LeagueID:    leagueID,
	}
	id, err := s.repo.CreateTeam(teamInfo)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s apiService) createStanding(i int, teamID uint, seasonID uint) error {
	standing := models.Standing{
		Rank:     uint(i),
		Form:     "",
		TeamID:   teamID,
		SeasonID: seasonID,
	}
	err := s.repo.CreateTables(standing)
	if err != nil {
		return err
	}
	return nil
}
func (s apiService) CreateStandings(info Info) error {
	var wg sync.WaitGroup
	errChan := make(chan error)
	reserve := make(chan struct{}, 4)

	leagueID, seasonID, err := s.getLeagueAndSeason(info)
	if err != nil {
		return err
	}

	standings, err := s.api.GetStandings(info.League, info.Season)
	if err != nil {
		return err
	}

	for i, v := range standings.Response[0].League.Standings[0] {
		wg.Add(1)
		go func(i int, v Standing) {
			defer wg.Done()
			reserve <- struct{}{}
			defer func() { <-reserve }()
			teamID, err := s.goFirstOrCreateTeam(v.Team.Name, uint(v.Team.ID), leagueID, seasonID)
			if err != nil {
				errChan <- err
				return
			}
			err = s.createStanding(i, teamID, seasonID)
			if err != nil {
				errChan <- err
				return
			}
		}(i, v)
	}
	wg.Wait()
	close(reserve)
	defer close(errChan)
	for err := range errChan {
		return err
	}
	return nil
}

// ________________________________________________________________
func (s apiService) createPlayer(player Player) (uint, error) {
	var playerID uint
	id, err := s.repo.FindTeam(player.Statistics[0].Team.Name)
	if err != nil {
		return 0, err
	}
	height, err := strconv.Atoi(player.Player.Height)
	if err != nil {
		return 0, err
	}
	weight, err := strconv.Atoi(player.Player.Weight)
	if err != nil {
		return 0, err
	}
	players := models.Player{
		Name:        player.Player.Name,
		Firstname:   player.Player.Firstname,
		Lastname:    player.Player.Lastname,
		Age:         uint(player.Player.Age),
		Nationality: player.Player.Nationality,
		Height:      uint(height),
		Weight:      uint(weight),
		Injuries:    player.Player.Injured,
		Photo:       player.Player.Photo,
		TeamID:      id,
	}
	playerID, err = s.repo.CreatePlayer(players)
	if err != nil {
		return 0, err
	}
	return playerID, nil
}
func (s apiService) createPlayerStatistics(playerID uint, player Player) error {
	res := player.Statistics[0]
	number, ok := res.Games.Number.(float64)
	if !ok {
		return errors.New("expected type float64")
	}
	statistics := models.PlayerStatistics{
		PlayerID: playerID,
		Number:   uint(number),
	}
	err := s.repo.CreatePlayerStatistics(playerID, statistics)
	if err != nil {
		return err
	}
	return nil
}
func (s apiService) createPlayerAndPlayerStatistics(player Players) error {
	for _, v := range player.Response {
		playerID, err := s.createPlayer(v)
		if err != nil {
			return err
		}
		err = s.createPlayerStatistics(playerID, v)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s apiService) CreatePlayers(info Info) error {
	var wg sync.WaitGroup
	// var m sync.Mutex
	// var once sync.Once
	errChan := make(chan error, 4)
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
					return
				}
				if page == 1 && totalPages == 1 {
					totalPages = players.Paging.Total
				}
				err = s.createPlayerAndPlayerStatistics(*players)
				if err != nil {
					errChan <- err
					return
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

// ________________________________________________________________
func (s apiService) CreateMatch(info Info) error {
	var wg sync.WaitGroup
	errChan := make(chan error, 4)
	reserve := make(chan struct{}, 4)
	round, err := s.repo.CountStandings(info.League, info.Season)
	if err != nil {
		return err
	}
	round2 := (round - 1) * 2
	for i := 1; i <= int(round2); i++ {
		wg.Add(1)
		go s.goCreateMatch(i, info, reserve, errChan, &wg)
	}
	wg.Wait()
	close(reserve)
	close(errChan)
	for err := range errChan {
		return err
	}
	return nil
}

func (s apiService) goCreateMatch(i int, info Info, reserve chan struct{}, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	reserve <- struct{}{}
	defer func() { <-reserve }()
	matchs, err := s.api.GetFixture(info.League, info.Season, i)
	if err != nil {
		errChan <- err
		return
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
		season, err := s.repo.FindSeason(uint(res.League.Season))
		if err != nil {
			errChan <- err
			return
		}
		match := models.Match{
			HomeTeamID: homeTeamID,
			AwayTeamID: awayTeamID,
			SeasonID:   season,
			Rounded:    uint(i),
			MatchDay:   res.Fixture.Date,
		}
		err = s.repo.CreateMatch(match)
		if err != nil {
			errChan <- err
			return
		}
	}

}
