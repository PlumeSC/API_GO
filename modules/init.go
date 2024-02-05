package modules

import (
	"false_api/modules/models"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*fiber.App, *gorm.DB) {
	if err := godotenv.Load(); err != nil {
		panic("error .env")
	}

	app := fiber.New()
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{}, &models.League{}, &models.Team{}, &models.Match{}, &models.Event{},
		&models.Standing{}, &models.Videos{}, &models.News{}, &models.Player{}, &models.Position{},
		&models.PlayerStatistics{})

	return app, db
}