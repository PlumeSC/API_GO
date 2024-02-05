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
	ID      uint `gorm:"primaryKey"`
	Name    string
	Country string
	Logo    string
	Flag    string
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
	LeagueID    string
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
}

type Standing struct {
	ID     uint `gorm:"primaryKey"`
	TeamID uint `gorm:"index"`
	Season uint `gorm:"index"`
	Rank   uint
	Played uint
	Won    uint
	Drawn  uint
	Lost   uint
	GF     uint
	GA     uint
	GD     uint
	Points uint
	Form   string
}

type Videos struct {
	ID       uint `gorm:"primaryKey"`
	MatchID  uint `gorm:"index"`
	TeamID   uint `gorm:"index"`
	Name     string
	Videos   string
	CreateAt time.Time
}

type Season struct {
	ID     uint `gorm:"primaryKey"`
	Season uint
}

type News struct {
	ID       uint `gorm:"primaryKey"`
	AdminID  uint `gorm:"index"`
	TeamID   uint `gorm:"index"`
	Title    string
	Content  string
	HeroImg  string
	CreateAt time.Time
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
	TeamID      string `gorm:"index"`
	Team        Team   `gorm:"foreignKey:TeamID"`
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
}
