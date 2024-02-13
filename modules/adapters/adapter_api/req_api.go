package adapterapi

import (
	"encoding/json"
	coreapi "false_api/modules/core/core_api"
	"fmt"
	"io"
	"net/http"
	"os"
)

type apiRequest struct{}

func NewApiRequest() *apiRequest {
	return &apiRequest{}
}

func (r apiRequest) GetTeam(id uint, league uint, season uint) (*coreapi.Team, error) {

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/teams?id=%v&league=%v&season=%v", id, league, season)
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

	var team coreapi.Team
	err = json.Unmarshal(body, &team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func (r apiRequest) GetLeague(league uint, season uint) (*coreapi.League, error) {

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/leagues?id=%v&season=%v", league, season)
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

	var leagueInfo coreapi.League
	err = json.Unmarshal(body, &leagueInfo)
	if err != nil {
		return nil, err
	}

	return &leagueInfo, nil
}
func (r apiRequest) GetStandings(league uint, season uint) (*coreapi.Standings, error) {

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

	var standingInfo coreapi.Standings
	err = json.Unmarshal(body, &standingInfo)
	if err != nil {
		return nil, err
	}

	return &standingInfo, nil
}

func (r apiRequest) GetPlayer(league uint, season uint, page int) (*coreapi.Players, error) {
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
	var playerInfo coreapi.Players
	err = json.Unmarshal(body, &playerInfo)
	if err != nil {
		return nil, err
	}
	return &playerInfo, nil
}

func (r apiRequest) GetFixture(league uint, season uint, round int) (*coreapi.Match, error) {
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?league=%v&season=%v&round=Regular%20Season%20-%20", league, season)
	url2 := fmt.Sprintf("%v%v", url, round)

	req, err := http.NewRequest("GET", url2, nil)
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
	var match coreapi.Match
	err = json.Unmarshal(body, &match)
	if err != nil {
		return nil, err
	}

	return &match, nil
}
