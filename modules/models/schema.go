package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Firstname  string
		Lastname   string
		Email      string `gorm:"unique"`
		Username   string `gorm:"unique"`
		Password   string
		ProfileImg string
		IsAdmin    bool
		IsVip      bool
		TeamID     uint
		Team       Team `gorm:"foreignKey:team_id"`
	}
	Team struct {
		gorm.Model
		Name        string `gorm:"unique"`
		CodeName    string `gorm:"unique"`
		Founded     uint
		Logo        string
		StadiumName string `gorm:"unique"`
		City        string
		Capacity    uint
		LeagueID    uint
		League      League   `gorm:"foreignKey:league_id"`
		Player      []Player `gorm:"foreignKey:TeamID"`
	}
	Player struct {
		gorm.Model
		Name        string
		Firstname   string
		Lastname    string
		Age         uint
		Nationality string
		Height      string
		Weight      string
		Injuries    bool
		Photo       string
		TeamID      uint
		Team        Team
	}
	PlayerStatistics struct {
		gorm.Model
		Appearances uint
		Lineup      uint
		Minutes     uint
		Number      uint
		Rating      uint
		Position    string
		PlayerID    uint
		Player      Player `gorm:"foreignKey:player_id"`
	}
	League struct {
		gorm.Model
		Name    string `gorm:"unique"`
		Country string
		Logo    string
		Flag    string
		ApiCode uint `gorm:"unique"`
	}
	Season struct {
		gorm.Model
		Season  uint `gorm:"unique"`
		Start   time.Time
		End     time.Time
		Current bool
	}
	LeagueSeason struct {
		gorm.Model
		LeagueID uint
		SeasonID uint
		League   League `gorm:"foreignKey:league_id"`
		Season   Season `gorm:"foreignKey:season_id"`
	}
	Standing struct {
		gorm.Model
		Rank           uint
		Played         uint `gorm:"default:0"`
		Won            uint `gorm:"default:0"`
		Drawn          uint `gorm:"default:0"`
		Lost           uint `gorm:"default:0"`
		GF             uint `gorm:"default:0"`
		GA             uint `gorm:"default:0"`
		GD             int  `gorm:"default:0"`
		Points         uint `gorm:"default:0"`
		Form           string
		LeagueSeasonID uint
		LeagueSeason   LeagueSeason `gorm:"foreignKey:league_season_id"`
		TeamID         uint
		Team           Team `gorm:"foreignKey:team_id"`
	}
	Match struct {
		gorm.Model
		Fixture        uint
		HomeGoal       uint `gorm:"default:0"`
		AwayGoal       uint `gorm:"default:0"`
		Rounded        uint
		MatchDay       time.Time
		MatchTime      string
		HomeTeamID     uint
		HomeTeam       Team `gorm:"foreignKey:home_team_id"`
		AwayTeamID     uint
		AwayTeam       Team `gorm:"foreignKey:away_team_id"`
		LeagueSeasonID uint
		LeagueSeason   LeagueSeason `gorm:"foreignKey:league_season_id"`
		Event          []Event      `gorm:"foreignKey:MatchID"`
	}
	Event struct {
		gorm.Model
		Assist    string
		EventTime uint
		Event     string
		PlayerID  uint
		Player    Player `gorm:"foreignKey:player_id"`
		MatchID   uint
		Match     Match `gorm:"foreignKey:MatchID"`
	}

	Videos struct {
		gorm.Model
		Name   string
		Videos string
		TeamID uint
		Team   Team `gorm:"foreignKey:team_id"`
	}
	News struct {
		gorm.Model
		Title   string
		Content string
		HeroImg string
		TeamID  uint
		Team    Team `gorm:"foreignKey:team_id"`
		UserID  uint
		User    User `gorm:"foreignKey:user_id"`
	}
)
