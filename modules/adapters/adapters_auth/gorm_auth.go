package adaptersauth

import (
	"false_api/modules/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepositoryImpl {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) WhereTeamID(team string) (uint, error) {
	teamID := models.Team{}
	if err := r.db.Where("name = ?", team).First(&teamID).Error; err != nil {
		return 0, err
	}
	return teamID.ID, nil
}

func (r *authRepositoryImpl) HashFunc(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func (r *authRepositoryImpl) CreateUser(user models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *authRepositoryImpl) WhereTeamName(id uint) (string, error) {
	team := models.Team{}
	if err := r.db.Where("ID = ?", id).First(&team).Error; err != nil {
		return "", err
	}
	return team.Name, nil
}
