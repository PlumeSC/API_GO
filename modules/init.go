package modules

import (
	"false_api/modules/models"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() (*fiber.App, *gorm.DB) {
	if err := godotenv.Load(); err != nil {
		panic("error .env")
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	_ = newLogger

	app := fiber.New()
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{}, &models.League{}, &models.Team{}, &models.Match{}, &models.Event{},
		&models.Standing{}, &models.Videos{}, &models.News{}, &models.Player{},
		&models.PlayerStatistics{})

	return app, db
}

func Middleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func Admin(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	splitToken := strings.Split(tokenString, "Bearer ")
	tokenString = splitToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "Admin" {
		return c.Status(fiber.StatusUnauthorized).SendString("not admin")
	}
	return c.Next()
}
