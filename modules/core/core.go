package core

import "time"

type Info struct {
	League uint `json:"league"`
	Season uint `json:"season"`
	Round  uint `json:"round"`
}

type League struct {
	Get        string `json:"get"`
	Parameters struct {
		ID     string `json:"id"`
		Season string `json:"season"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		League struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			Logo string `json:"logo"`
		} `json:"league"`
		Country struct {
			Name string `json:"name"`
			Code string `json:"code"`
			Flag string `json:"flag"`
		} `json:"country"`
		Seasons []struct {
			Year     int    `json:"year"`
			Start    string `json:"start"`
			End      string `json:"end"`
			Current  bool   `json:"current"`
			Coverage struct {
				Fixtures struct {
					Events             bool `json:"events"`
					Lineups            bool `json:"lineups"`
					StatisticsFixtures bool `json:"statistics_fixtures"`
					StatisticsPlayers  bool `json:"statistics_players"`
				} `json:"fixtures"`
				Standings   bool `json:"standings"`
				Players     bool `json:"players"`
				TopScorers  bool `json:"top_scorers"`
				TopAssists  bool `json:"top_assists"`
				TopCards    bool `json:"top_cards"`
				Injuries    bool `json:"injuries"`
				Predictions bool `json:"predictions"`
				Odds        bool `json:"odds"`
			} `json:"coverage"`
		} `json:"seasons"`
	} `json:"response"`
}

type Standings struct {
	Get        string `json:"get"`
	Parameters struct {
		League string `json:"league"`
		Season string `json:"season"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		League struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Country   string `json:"country"`
			Logo      string `json:"logo"`
			Flag      string `json:"flag"`
			Season    int    `json:"season"`
			Standings [][]struct {
				Rank int `json:"rank"`
				Team struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
					Logo string `json:"logo"`
				} `json:"team"`
				Points      int    `json:"points"`
				GoalsDiff   int    `json:"goalsDiff"`
				Group       string `json:"group"`
				Form        string `json:"form"`
				Status      string `json:"status"`
				Description string `json:"description"`
				All         struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"all"`
				Home struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"home"`
				Away struct {
					Played int `json:"played"`
					Win    int `json:"win"`
					Draw   int `json:"draw"`
					Lose   int `json:"lose"`
					Goals  struct {
						For     int `json:"for"`
						Against int `json:"against"`
					} `json:"goals"`
				} `json:"away"`
				Update time.Time `json:"update"`
			} `json:"standings"`
		} `json:"league"`
	} `json:"response"`
}

type Standing struct {
	Rank int `json:"rank"`
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"team"`
	Points      int    `json:"points"`
	GoalsDiff   int    `json:"goalsDiff"`
	Group       string `json:"group"`
	Form        string `json:"form"`
	Status      string `json:"status"`
	Description string `json:"description"`
	All         struct {
		Played int `json:"played"`
		Win    int `json:"win"`
		Draw   int `json:"draw"`
		Lose   int `json:"lose"`
		Goals  struct {
			For     int `json:"for"`
			Against int `json:"against"`
		} `json:"goals"`
	} `json:"all"`
	Home struct {
		Played int `json:"played"`
		Win    int `json:"win"`
		Draw   int `json:"draw"`
		Lose   int `json:"lose"`
		Goals  struct {
			For     int `json:"for"`
			Against int `json:"against"`
		} `json:"goals"`
	} `json:"home"`
	Away struct {
		Played int `json:"played"`
		Win    int `json:"win"`
		Draw   int `json:"draw"`
		Lose   int `json:"lose"`
		Goals  struct {
			For     int `json:"for"`
			Against int `json:"against"`
		} `json:"goals"`
	} `json:"away"`
	Update time.Time `json:"update"`
}

type Team struct {
	Get        string `json:"get"`
	Parameters struct {
		League string `json:"league"`
		ID     string `json:"id"`
		Season string `json:"season"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Team struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Code     string `json:"code"`
			Country  string `json:"country"`
			Founded  int    `json:"founded"`
			National bool   `json:"national"`
			Logo     string `json:"logo"`
		} `json:"team"`
		Venue struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Address  string `json:"address"`
			City     string `json:"city"`
			Capacity int    `json:"capacity"`
			Surface  string `json:"surface"`
			Image    string `json:"image"`
		} `json:"venue"`
	} `json:"response"`
}

type Players struct {
	Get        string `json:"get"`
	Parameters struct {
		League string `json:"league"`
		Page   string `json:"page"`
		Season string `json:"season"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Player struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
			Age       int    `json:"age"`
			Birth     struct {
				Date    string `json:"date"`
				Place   string `json:"place"`
				Country string `json:"country"`
			} `json:"birth"`
			Nationality string `json:"nationality"`
			Height      string `json:"height"`
			Weight      string `json:"weight"`
			Injured     bool   `json:"injured"`
			Photo       string `json:"photo"`
		} `json:"player"`
		Statistics []struct {
			Team struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Logo string `json:"logo"`
			} `json:"team"`
			League struct {
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Country string `json:"country"`
				Logo    string `json:"logo"`
				Flag    string `json:"flag"`
				Season  int    `json:"season"`
			} `json:"league"`
			Games struct {
				Appearences interface{} `json:"appearences"`
				Lineups     interface{} `json:"lineups"`
				Minutes     interface{} `json:"minutes"`
				Number      interface{} `json:"number"`
				Position    string      `json:"position"`
				Rating      interface{} `json:"rating"`
				Captain     bool        `json:"captain"`
			} `json:"games"`
			Substitutes struct {
				In    interface{} `json:"in"`
				Out   interface{} `json:"out"`
				Bench interface{} `json:"bench"`
			} `json:"substitutes"`
			Shots struct {
				Total interface{} `json:"total"`
				On    interface{} `json:"on"`
			} `json:"shots"`
			Goals struct {
				Total    interface{} `json:"total"`
				Conceded interface{} `json:"conceded"`
				Assists  interface{} `json:"assists"`
				Saves    interface{} `json:"saves"`
			} `json:"goals"`
			Passes struct {
				Total    interface{} `json:"total"`
				Key      interface{} `json:"key"`
				Accuracy interface{} `json:"accuracy"`
			} `json:"passes"`
			Tackles struct {
				Total         interface{} `json:"total"`
				Blocks        interface{} `json:"blocks"`
				Interceptions interface{} `json:"interceptions"`
			} `json:"tackles"`
			Duels struct {
				Total interface{} `json:"total"`
				Won   interface{} `json:"won"`
			} `json:"duels"`
			Dribbles struct {
				Attempts interface{} `json:"attempts"`
				Success  interface{} `json:"success"`
				Past     interface{} `json:"past"`
			} `json:"dribbles"`
			Fouls struct {
				Drawn     interface{} `json:"drawn"`
				Committed interface{} `json:"committed"`
			} `json:"fouls"`
			Cards struct {
				Yellow    interface{} `json:"yellow"`
				Yellowred interface{} `json:"yellowred"`
				Red       interface{} `json:"red"`
			} `json:"cards"`
			Penalty struct {
				Won      interface{} `json:"won"`
				Commited interface{} `json:"commited"`
				Scored   interface{} `json:"scored"`
				Missed   interface{} `json:"missed"`
				Saved    interface{} `json:"saved"`
			} `json:"penalty"`
		} `json:"statistics"`
	} `json:"response"`
}

type Match struct {
	Get        string `json:"get"`
	Parameters struct {
		League string `json:"league"`
		Round  string `json:"round"`
		Season string `json:"season"`
	} `json:"parameters"`
	Errors  []interface{} `json:"errors"`
	Results int           `json:"results"`
	Paging  struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Response []struct {
		Fixture struct {
			ID        int       `json:"id"`
			Referee   string    `json:"referee"`
			Timezone  string    `json:"timezone"`
			Date      time.Time `json:"date"`
			Timestamp int       `json:"timestamp"`
			Periods   struct {
				First  int `json:"first"`
				Second int `json:"second"`
			} `json:"periods"`
			Venue struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				City string `json:"city"`
			} `json:"venue"`
			Status struct {
				Long    string `json:"long"`
				Short   string `json:"short"`
				Elapsed int    `json:"elapsed"`
			} `json:"status"`
		} `json:"fixture"`
		League struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Country string `json:"country"`
			Logo    string `json:"logo"`
			Flag    string `json:"flag"`
			Season  int    `json:"season"`
			Round   string `json:"round"`
		} `json:"league"`
		Teams struct {
			Home struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Winner bool   `json:"winner"`
			} `json:"home"`
			Away struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Logo   string `json:"logo"`
				Winner bool   `json:"winner"`
			} `json:"away"`
		} `json:"teams"`
		Goals struct {
			Home int `json:"home"`
			Away int `json:"away"`
		} `json:"goals"`
		Score struct {
			Halftime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"halftime"`
			Fulltime struct {
				Home int `json:"home"`
				Away int `json:"away"`
			} `json:"fulltime"`
			Extratime struct {
				Home interface{} `json:"home"`
				Away interface{} `json:"away"`
			} `json:"extratime"`
			Penalty struct {
				Home interface{} `json:"home"`
				Away interface{} `json:"away"`
			} `json:"penalty"`
		} `json:"score"`
	} `json:"response"`
}
