package competitivecore

import (
	"false_api/modules/core"
	matchscore "false_api/modules/core/matchs_core"
	seasoncore "false_api/modules/core/season_core"
	"false_api/modules/models"
	"fmt"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
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

	isCron := cron.New(cron.WithLocation(loc))

	isCron.Start()
	go func() {
		s.newDay()
	}()

	_, err = isCron.AddFunc("@daily", s.newDay)
	if err != nil {
		return err
	}
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
	cron := cron.New(cron.WithLocation(loc))
	match := map[time.Time]bool{}
	var matches []time.Time

	for _, v := range data {
		if _, ok := match[v.MatchDay]; !ok {
			match[v.MatchDay] = true
			matches = append(matches, v.MatchDay)
		}
	}
	for _, v := range matches {
		matchesHour := v.Hour()
		matchesMinute := v.Minute()

		convertCron := s.convertTimeTocron(v)
		_ = convertCron
		fmt.Println(time.Now())
		convertCron = "0 20 3 12 3"
		fmt.Println(convertCron)

		_, err := cron.AddFunc(convertCron, func() {
			fmt.Println("xx")
			go func(hour int, minute int) {
				currentTime := time.Now()
				currentTimeHour := currentTime.Hour()
				currentTimeMinute := currentTime.Minute()

				if currentTimeHour == hour && currentTimeMinute == minute {
					for {
						status := s.liveScore(data[0].League, data[0].Season, data[0].LeagueSeasonID)
						if !status {
							return
						}
						time.Sleep(1 * time.Minute)
					}
				}
			}(matchesHour, matchesMinute)
		})
		cron.Start()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (s compService) convertTimeTocron(time time.Time) string {
	return fmt.Sprintf("0 %d %d %d %d ", time.Minute(), time.Hour(), time.Day(), time.Month())
}

func (s compService) liveScore(league uint, season uint, lsID uint) bool {
	matches, err := s.api.GetLiveScore(league, season)
	if err != nil {
		fmt.Println(err)
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
		if v.Fixture.Status.Long == "Match Finished" {
			return false
		}
	}
	return true
}
