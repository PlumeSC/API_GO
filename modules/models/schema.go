package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname  string `gorm:"not null"`
	Lastname   string `gorm:"not null"`
	Email      string `gorm:"type:varchar(100);unique_index"`
	Username   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	ProfileImg string
	TeamID     uint `gorm:"index"`
	IsAdmin    bool `gorm:"default:false"`
	IsVip      bool `gorm:"default:false"`
	Team       Team `gorm:"foreignKey:team_id"`
}

type League struct {

	gorm.Model
	Name    string `gorm:"unique"`
	Country string
	Logo    string
	Flag    string
	ApiCode uint `gorm:"unique"`
}

type Season struct {
	gorm.Model
	Season  uint
	Start   time.Time
	End     time.Time
	Current bool
}

type LeagueSeason struct {
	ID       uint `gorm:"primaryKey"`
	SeasonID uint
	Season   Season `gorm:"foreignKey:season_id"`
	LeagueID uint
	League   League `gorm:"foreignKey:league_id"`

}

type Team struct {
	gorm.Model
	Name        string `gorm:"unique"`
	CodeName    string `gorm:"unique"`
	Founded     uint
	Logo        string
	StadiumName string `gorm:"unique"`
	City        string
	Capacity    uint
	LeagueID    uint
	League      League `gorm:"foreignKey:league_id"`
}

type Match struct {
	gorm.Model
	HomeTeamID     uint `gorm:"index"`
	AwayTeamID     uint `gorm:"index"`
	Fixture        uint
	HomeGoal       uint `gorm:"default:0"`
	AwayGoal       uint `gorm:"default:0"`
	Rounded        uint
	MatchDay       time.Time
	MatchTime      string
	LeagueSeasonID uint         `gorm:"index"`
	LeagueSeason   LeagueSeason `gorm:"foreignKey:league_season_id"`
	HomeTeam       Team         `gorm:"foreignKey:home_team_id"`
	AwayTeam       Team         `gorm:"foreignKey:away_team_id"`

}

type Event struct {
	ID        int  `gorm:"primaryKey"`
	MatchID   uint `gorm:"index"`
	TeamID    uint `gorm:"index"`
	Player    string
	Assist    string
	EventTime uint
	Event     string
	Team      Team  `gorm:"foreignKey:TeamID"`
	Match     Match `gorm:"foreignKey:MatchID"`
}

type Standing struct {
	gorm.Model
	TeamID         uint `gorm:"index"`
	LeagueSeasonID uint `gorm:"index"`
	Rank           uint
	Played         uint `gorm:"default:0"`
	Won            uint `gorm:"default:0"`
	Drawn          uint `gorm:"default:0"`
	Lost           uint `gorm:"default:0"`
	GF             uint `gorm:"default:0"`
	GA             uint `gorm:"default:0"`
	GD             uint `gorm:"default:0"`
	Points         uint `gorm:"default:0"`
	Form           string
	Team           Team         `gorm:"foreignKey:team_id"`
	LeagueSeason   LeagueSeason `gorm:"foreignKey:league_season_id"`
}

type Videos struct {
	gorm.Model
	MatchID uint `gorm:"index"`
	TeamID  uint `gorm:"index"`
	Name    string
	Videos  string
	Team    Team  `gorm:"foreignKey:team_id"`
	Match   Match `gorm:"foreignKey:match_id"`

}

type News struct {
	gorm.Model
	UserID   uint `gorm:"index"`
	TeamID   uint `gorm:"index"`
	Title    string
	Content  string
	HeroImg  string
	CreateAt time.Time

	Team     Team `gorm:"foreignKey:team_id"`
	User     User `gorm:"foreignKey:user_id"`

}

type Player struct {
	gorm.Model
	Name        string
	Firstname   string
	Lastname    string
	Age         uint
	Nationality string
	Height      uint
	Weight      uint
	Injuries    bool
	Photo       string

	TeamID uint `gorm:"index"`

	Team   Team `gorm:"foreignKey:team_id"`

}

type PlayerStatistics struct {
	ID          uint `gorm:"primaryKey"`
	PlayerID    uint `gorm:"index"`
	Appearances uint
	Lineup      uint
	Minutes     uint
	Number      uint
	Rating      uint
	Position    string
	Player      Player `gorm:"foreignKey:player_id"`
}
