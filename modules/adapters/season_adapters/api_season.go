package seasonadapters

import (
	"encoding/json"
	core "false_api/modules/core"
	"fmt"
	"io"
	"net/http"
	"os"
)

type apiFootball struct{}

func NewApiFootball() *apiFootball {
	return &apiFootball{}
}

func (r apiFootball) GetLeague(leagueApi uint, season uint) (*core.League, error) {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/leagues?id=%v&season=%v", leagueApi, season)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_FOOTBALL"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var leagueInfo core.League
	err = json.Unmarshal(body, &leagueInfo)
	if err != nil {
		return nil, err
	}
	return &leagueInfo, nil
}

func (r apiFootball) GetStandings(league uint, season uint) (*core.Standings, error) {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/standings?season=%v&league=%v", season, league)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_FOOTBALL"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var standingsInfo core.Standings
	err = json.Unmarshal(body, &standingsInfo)
	if err != nil {
		return nil, err
	}
	return &standingsInfo, nil
}

func (r apiFootball) GetTeam(codeTeam uint, codeLeague uint, season uint) (*core.Team, error) {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/teams?id=%v&league=%v&season=%v", codeTeam, codeLeague, season)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_FOOTBALL"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var teamInfo core.Team
	err = json.Unmarshal(body, &teamInfo)
	if err != nil {
		return nil, err
	}
	return &teamInfo, nil
}

func (r apiFootball) GetPlayer(league uint, season uint, page int) (*core.Players, error) {

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/players?league=%v&season=%v&page=%v", league, season, page)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_FOOTBALL"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var playerInfo core.Players
	err = json.Unmarshal(body, &playerInfo)
	if err != nil {
		return nil, err
	}
	return &playerInfo, nil
}

func (r apiFootball) GetFixture(league uint, season uint, round int) (*core.Match, error) {
	x := "%20"
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?league=%v&season=%v&round=Regular%vSeason%v-%v%v", league, season, x, x, x, round)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_FOOTBALL"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var match core.Match
	err = json.Unmarshal(body, &match)
	if err != nil {
		return nil, err
	}

	return &match, nil
}
