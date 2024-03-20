package matchsadapters

import (
	matchscore "false_api/modules/core/matchs_core"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type matchsHandler struct {
	service matchscore.MatchsService
}

func NewMatchsHandler(service matchscore.MatchsService) *matchsHandler {
	return &matchsHandler{service: service}
}

func (h matchsHandler) GetAll(c *fiber.Ctx) error {
	teamNameParam := c.Query("team")
	roundParam := c.Query("round")
	api_codeParam := c.Query("api_code")
	seasonParam := c.Query("season")
	var trimmedStr string
	if teamNameParam != "" {
		runes := []rune(teamNameParam)
		trimmedStr = string(runes[1 : len(runes)-1])
	}
	var round int
	var err error
	if roundParam != "" {
		round, err = strconv.Atoi(roundParam)
		if err != nil {
			fmt.Println("round", err)
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	}
	season, err := strconv.Atoi(seasonParam)
	if err != nil {
		fmt.Println("season")
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	api_code, err := strconv.Atoi(api_codeParam)
	if err != nil {
		fmt.Println("api")
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	params := map[string]interface{}{
		"team_name": trimmedStr,
		"round":     uint(round),
		"api_code":  uint(api_code),
		"season":    uint(season),
	}
	matches, err := h.service.GetMatchs(params)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(matches)
}

func (h matchsHandler) UpdateMatch(c *fiber.Ctx) error {
	apiParam := c.Query("api_code")
	seasonParam := c.Query("season")
	roundParam := c.Query("round")

	api, err := strconv.Atoi(apiParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	season, err := strconv.Atoi(seasonParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	round, err := strconv.Atoi(roundParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	data := map[string]int{
		"api_code": api,
		"season":   season,
		"round":    round,
	}
	err = h.service.UpdateMatchs(data)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"msg": "ok",
	})
}

func (h matchsHandler) GetPlayer(c *fiber.Ctx) error {
	name := c.Query("playername")
	player, err := h.service.GetPlayer(name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(player)

}
