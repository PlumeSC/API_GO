package coreauth

import (
	"false_api/modules/models"
)

type AuthRepository interface {
	CreateUser(models.User) (*models.User, error)
	HashFunc(string) (string, error)
	WhereTeamID(string) (uint, error)
	WhereTeamName(uint) (string, error)
	CheckUser(string) (*models.User, error)
}
