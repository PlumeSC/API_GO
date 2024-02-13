package coreapi

import (
	"errors"
	"false_api/modules/models"
	"strconv"
	"time"
)

type ApiService interface {
	CreateTables(Info) error
	CreatePlayer(Info) error
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

func (s apiService) CreateLeague(league League) (*uint, error) {
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
	if err == nil {
		return leagueID, err
	}
	return leagueID, nil
}

func (s apiService) FindOrCreateSeason(league League, leagueID uint) (uint, error) {
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
		LeagueID: leagueID,
		Start:    start,
		End:      end,
		Current:  league.Response[0].Seasons[0].Current,
	}
	seasonID, err := s.repo.FindOrCreateSeason(season)
	if err != nil {
		return 0, err
	}
	return seasonID, nil
}

func (s apiService) CreateSeason(league League, leagueID *uint) (*uint, error) {
	season := models.Season{}
	layout := "2006-01-02"
	start, err := time.Parse(layout, league.Response[0].Seasons[0].Start)
	if err != nil {
		return &season.ID, err
	}
	end, err := time.Parse(layout, league.Response[0].Seasons[0].End)
	if err != nil {
		return &season.ID, err
	}
	season = models.Season{
		Season:   uint(league.Response[0].Seasons[0].Year),
		LeagueID: *leagueID,
		Start:    start,
		End:      end,
		Current:  league.Response[0].Seasons[0].Current,
	}
	return &season.ID, nil
}

func (s apiService) CreateTables(info Info) error {
	league := &League{}
	var leagueID *uint
	var seasonID *uint
	var teamID *uint
	haveSeason, seasonID, err := s.repo.CheckSeason(info.Season)
	if err != nil {
		return err
	}
	if !haveSeason {
		var haveLeague bool
		haveLeague, leagueID, err = s.repo.CheckLeague(info.League)
		if err != nil {
			return err
		}
		if !haveLeague {
			league, err = s.api.GetLeague(info.League, info.Season)
			if err != nil {
				return err
			}
			leagueID, err = s.CreateLeague(*league)
			if err != nil {
				return err
			}
		}
		seasonID, err = s.CreateSeason(*league, leagueID)
		if err != nil {
			return err
		}
	}

	standings, err := s.api.GetStandings(info.League, info.Season)
	if err != nil {
		return err
	}

	for i, v := range standings.Response[0].League.Standings[0] {
		teamID, err = s.repo.FindTeam(v.Team.Name)
		if err != nil {
			return err
		}
		if *teamID == 0 {
			team, err := s.api.GetTeam(uint(v.Team.ID), *leagueID, *seasonID)
			if err != nil {
				return err
			}
			res := team.Response[0]
			teamInfo := models.Team{
				Name:        res.Team.Name,
				CodeName:    res.Team.Code,
				Founded:     uint(res.Team.Founded),
				Logo:        res.Team.Logo,
				StadiumName: res.Venue.Name,
				City:        res.Venue.City,
				Capacity:    uint(res.Venue.Capacity),
				LeagueID:    *leagueID,
			}
			teamID, err = s.repo.CreateTeam(teamInfo)
			if err != nil {
				return err
			}
		}
		standing := models.Standing{
			Rank:     uint(i),
			Form:     "",
			TeamID:   *teamID,
			SeasonID: *seasonID,
		}
		err = s.repo.CreateTables(standing)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s apiService) Player(player Player) error {
	// players := models.Player{}
	for _, v := range player.Response {
		id, err := s.repo.FindTeam(v.Statistics[0].Team.Name)
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
		players := models.Player{
			Name:        v.Player.Name,
			Firstname:   v.Player.Firstname,
			Lastname:    v.Player.Lastname,
			Age:         uint(v.Player.Age),
			Nationality: v.Player.Nationality,
			Height:      uint(height),
			Weight:      uint(weight),
			Injuries:    v.Player.Injured,
			Photo:       v.Player.Photo,
			TeamID:      *id,
		}
		err = s.repo.CreatePlayer(players)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s apiService) CreatePlayer(info Info) error {
	pages := 1
	totalPages := 1
	for pages <= totalPages {
		players, err := s.api.GetPlayer(info.League, info.Season, pages)
		if err != nil {
			return err
		}
		if err = s.Player(*players); err != nil {
			return err
		}
		totalPages = players.Paging.Total
		pages++
	}
	return nil
}

func (s apiService) Match(match Match) {

}

func (s apiService) CreateMatch(info Info) error {
	if info.Round > 38 {
		return errors.New("round must equal or less than 38")
	}
	matchs, err := s.api.GetFixture(info.League, info.Season, int(info.Round))
	if err != nil {
		return err
	}
	for _, v := range matchs.Response {
		homeTeamID, err := s.repo.FindTeam(v.Teams.Home.Name)
		if err != nil {
			return err
		}
		awayTeamID, err := s.repo.FindTeam(v.Teams.Away.Name)
		if err != nil {
			return err
		}
		_, seasonID, err := s.repo.CheckSeason(info.Season)
		if err != nil {
			return err
		}
		match := models.Match{
			HomeTeamID: *homeTeamID,
			AwayTeamID: *awayTeamID,
			SeasonID:   *seasonID,
			HomeGoal:   uint(v.Goals.Home),
			AwayGoal:   uint(v.Goals.Away),
			Rounded:    uint(info.Round),
			MatchDay:   v.Fixture.Date,
		}
		err = s.repo.CreateMatch(*homeTeamID, *awayTeamID, match)
		if err != nil {
			return err
		}
	}

	return nil
}
