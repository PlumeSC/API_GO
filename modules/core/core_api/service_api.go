package coreapi

import (
	"false_api/modules/models"
	"time"
)

type ApiService interface {
	CreateLeague(league League) (*uint, error)
	FindOrCreateSeason(league League, leagueID uint) (uint, error)
	CreateTables(StandingInfo) error
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

func (s apiService) CreateTables(info StandingInfo) error {
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
