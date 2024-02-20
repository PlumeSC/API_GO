package authcore

import (
	"false_api/modules/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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
	Login(Login) (string, *UserInfo, error)
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
	token, userInfo, err := s.prepareInfo(*userx)
	if err != nil {
		return "", nil, err
	}
	return token, userInfo, nil
}

func (s *authServiceImpl) Login(userLogin Login) (string, *UserInfo, error) {
	user, err := s.repo.CheckUser(userLogin.Username)
	if err != nil {
		return "", nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))
	if err != nil {
		return "", nil, err
	}
	token, userInfo, err := s.prepareInfo(*user)
	if err != nil {
		return "", nil, err
	}
	return token, userInfo, nil
}

func (s *authServiceImpl) prepareInfo(user models.User) (string, *UserInfo, error) {
	token, err := s.createToken(user)
	if err != nil {
		return "", nil, err
	}
	teamName, err := s.repo.WhereTeamName(user.TeamID)
	if err != nil {
		return "", nil, err
	}
	userInfo := UserInfo{
		Firstname:  user.Username,
		Lastname:   user.Lastname,
		ProfileImg: user.ProfileImg,
		TeamName:   teamName,
		IsAdmin:    user.IsAdmin,
		IsVip:      user.IsVip,
	}
	return token, &userInfo, nil
}

func (s *authServiceImpl) createToken(user models.User) (string, error) {
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
