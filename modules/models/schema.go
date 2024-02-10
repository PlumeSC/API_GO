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
	Team       Team `gorm:"foreignKey:TeamID"`
}

type League struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	LeagueID uint
	Country  string
	Code     string
	Logo     string
	Flag     string
}

type Season struct {
	ID       uint `gorm:"primaryKey"`
	Season   uint
	LeagueID uint
	League   League `gorm:"foreignKey:LeagueID"`
	Start    time.Time
	End      time.Time
	Current  bool
}

type Team struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique"`
	CodeName    string `gorm:"unique"`
	Founded     uint
	Logo        string
	StadiumName string `gorm:"unique"`
	City        string
	Capacity    uint
	LeagueID    uint
	League      League `gorm:"foreignKey:LeagueID"`
}

type Match struct {
	ID        uint `gorm:"primaryKey"`
	HomeTeam  uint `gorm:"index"`
	AwayTeam  uint `gorm:"index"`
	Fixture   uint
	HomeGoal  uint
	AwayGoal  uint
	Season    uint `gorm:"index"`
	Rounded   uint
	MatchDay  time.Time
	MatchTime string
	SeasonRef Season `gorm:"foreignKey:Season"`
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
	ID       uint `gorm:"primaryKey"`
	TeamID   uint `gorm:"index"`
	SeasonID uint `gorm:"index"`
	Rank     uint
	Played   uint `gorm:"default:0"`
	Won      uint `gorm:"default:0"`
	Drawn    uint `gorm:"default:0"`
	Lost     uint `gorm:"default:0"`
	GF       uint `gorm:"default:0"`
	GA       uint `gorm:"default:0"`
	GD       uint `gorm:"default:0"`
	Points   uint `gorm:"default:0"`
	Form     string
	Team     Team   `gorm:"foreignKey:TeamID"`
	Season   Season `gorm:"foreignKey:SeasonID"`
}

type Videos struct {
	ID       uint `gorm:"primaryKey"`
	MatchID  uint `gorm:"index"`
	TeamID   uint `gorm:"index"`
	Name     string
	Videos   string
	CreateAt time.Time
	Team     Team  `gorm:"foreignKey:TeamID"`
	Match    Match `gorm:"foreignKey:MatchID"`
}

type News struct {
	ID       uint `gorm:"primaryKey"`
	AdminID  uint `gorm:"index"`
	TeamID   uint `gorm:"index"`
	Title    string
	Content  string
	HeroImg  string
	CreateAt time.Time
	Team     Team `gorm:"foreignKey:TeamID"`
	User     User `gorm:"foreignKey:AdminID"`
}

type Player struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Firstname   string
	Lastname    string
	Age         uint
	Nationality string
	Height      uint
	Weight      uint
	Injuries    bool
	Photo       string
	TeamID      uint `gorm:"index"`
	Team        Team `gorm:"foreignKey:TeamID"`
}

type Position struct {
	ID           uint `gorm:"primaryKey"`
	PositionName string
	Position     string
}

type PlayerStatistics struct {
	ID          uint `gorm:"primaryKey"`
	PlayerID    uint `gorm:"index"`
	Appearances uint
	Lineup      uint
	Minutes     uint
	Number      uint
	PositionID  uint     `gorm:"index"`
	Position    Position `gorm:"foreignKey:PositionID"`
	Player      Player   `gorm:"foreignKey:PlayerID"`
}
