package coreauth

import (
	"false_api/modules/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Register struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Team      string `json:"team"`
}
type UserInfo struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	ProfileImg string `json:"img"`
	TeamName   string `json:"team"`
	IsAdmin    bool   `json:"isAdmin"`
	IsVip      bool   `json:"isVip"`
}

type AuthService interface {
	Register(Register) (string, *UserInfo, error)
	CreateToken(models.User) (string, error)
	// Login(Login)
}

type authServiceImpl struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) *authServiceImpl {
	return &authServiceImpl{repo: repo}
}

func (s *authServiceImpl) Register(userReg Register) (string, *UserInfo, error) {
	teamID, err := s.repo.WhereTeamID(userReg.Team)
	if err != nil {
		return "", nil, err
	}
	userReg.Password, err = s.repo.HashFunc(userReg.Password)
	if err != nil {
		return "", nil, err
	}
	user := models.User{
		Firstname: userReg.Firstname,
		Lastname:  userReg.Lastname,
		Email:     userReg.Email,
		Username:  userReg.Username,
		Password:  userReg.Password,
		TeamID:    teamID,
	}
	userx, err := s.repo.CreateUser(user)
	if err != nil {
		return "", nil, err
	}
	token, err := s.CreateToken(*userx)
	if err != nil {
		return "", nil, err
	}
	teamName, err := s.repo.WhereTeamName(user.TeamID)
	if err != nil {
		return "", nil, err
	}
	userInfo := UserInfo{
		Firstname:  userx.Username,
		Lastname:   userx.Lastname,
		ProfileImg: userx.ProfileImg,
		TeamName:   teamName,
		IsAdmin:    userx.IsAdmin,
		IsVip:      userx.IsVip,
	}
	return token, &userInfo, nil
}

func (s *authServiceImpl) CreateToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return t, nil
}
