package competitivecore

import (
	"false_api/modules/core"
	matchscore "false_api/modules/core/matchs_core"
	seasoncore "false_api/modules/core/season_core"
	"false_api/modules/models"
	"fmt"
	"strconv"
	"time"

	"github.com/robfig/cron"
)

type LiveData struct {
	ID             uint
	Round          uint
	MatchDay       time.Time
	LeagueSeasonID uint
	League         uint
	Season         uint
}

type CompService interface {
	Comp(core.Info) error
}

type compService struct {
	repo        CompRepository
	api         seasoncore.SeasonApi
	repoMatches matchscore.MatchsRepository
}

func NewCompService(repo CompRepository, api seasoncore.SeasonApi, repoMatches matchscore.MatchsRepository) *compService {
	return &compService{
		repo:        repo,
		api:         api,
		repoMatches: repoMatches,
	}
}

func (s compService) Comp(info core.Info) error {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	isCron := cron.NewWithLocation(loc)
	err = isCron.AddFunc("0 0 0 * * *", s.newDay)
	if err != nil {
		return err
	}

	data, err := s.getData()
	if err != nil {
		fmt.Println(err)
	}
	s.Live(data)

	return nil
}

func (s compService) newDay() {
	data, err := s.getData()
	if err != nil {
		fmt.Println(err)
	}
	s.Live(data)
}

func (s compService) getData() ([]LiveData, error) {
	var data []LiveData
	matches, err := s.repo.GetAllToday()
	if err != nil {
		return nil, err
	}
	for _, v := range matches {
		data = append(data, LiveData{
			ID:             v.ID,
			Round:          v.Rounded,
			MatchDay:       v.MatchDay,
			LeagueSeasonID: v.LeagueSeasonID,
			League:         v.LeagueSeason.League.ApiCode,
			Season:         v.LeagueSeason.Season.Season,
		})
	}
	return data, nil
}

func (s compService) Live(data []LiveData) {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	cron := cron.NewWithLocation(loc)
	fmt.Println("open")
	matches := []time.Time{}
	for _, v := range data {
		matches = append(matches, v.MatchDay)
	}
	fmt.Println(matches)
	for _, v := range matches {
		matchesHour := v.Hour()
		matchesMinute := v.Minute()

		convertCron := s.convertTimeTocron(v)
		fmt.Println(convertCron)

		cron.AddFunc(convertCron, func() {
			fmt.Println("live1")
			go func(hour int, minute int) {
				fmt.Println("live2")
				currentTime := time.Now()
				currentTimeHour := currentTime.Hour()
				currentTimeMinute := currentTime.Minute()

				if currentTimeHour == hour && currentTimeMinute == minute {
					for {
						status := s.liveScore(data[0].League, data[0].Season, data[0].LeagueSeasonID)
						if status == "FT" {
							return
						}
						time.Sleep(1 * time.Minute)
					}
				}
			}(matchesHour, matchesMinute)
		})
	}
}

func (s compService) convertTimeTocron(time time.Time) string {
	cron := fmt.Sprintf("%d %d %d %d * %d", time.Minute(), time.Hour(), time.Day(), time.Month(), time.Year())
	return cron
}

func (s compService) liveScore(league uint, season uint, lsID uint) string {
	matches, err := s.api.GetLiveScore(league, season)
	if err != nil {
		fmt.Println(err)
	}

	if matches.Response[0].Fixture.Status.Short == "FT" {
		return "FT"
	}
	for _, v := range matches.Response {
		matchTime := strconv.Itoa(v.Fixture.Status.Elapsed)
		homeID, err := s.repoMatches.FindTeam(v.Teams.Home.Name)
		if err != nil {
			fmt.Println(err)
		}
		awayID, err := s.repoMatches.FindTeam(v.Teams.Away.Name)
		if err != nil {
			fmt.Println(err)
		}
		round := v.League.Round[17:]
		rounded, err := strconv.Atoi(round)
		if err != nil {
			fmt.Println(err)
		}
		data := models.Match{
			Fixture:        uint(v.Fixture.ID),
			HomeGoal:       uint(v.Goals.Home),
			AwayGoal:       uint(v.Goals.Away),
			Rounded:        uint(rounded),
			MatchDay:       v.Fixture.Date,
			MatchTime:      matchTime,
			HomeTeamID:     homeID,
			AwayTeamID:     awayID,
			LeagueSeasonID: lsID,
		}
		err = s.repoMatches.UpdateMatch(data, homeID, awayID, lsID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("update")
	}
	return "ok"
}
