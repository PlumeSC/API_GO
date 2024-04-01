package standingsadapters

import (
	"false_api/modules/core"
	standingscore "false_api/modules/core/standings_core"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type standingsHandler struct {
	service standingscore.StandingsService
}

func NewStandingsHandler(service standingscore.StandingsService) *standingsHandler {
	return &standingsHandler{service: service}
}

func (h standingsHandler) GetStandings(c *fiber.Ctx) error {
	seasonParam := c.Query("season")
	if seasonParam == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing 'season' query parameter")
	}

	leagueParam := c.Query("league")
	if leagueParam == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing 'league' query parameter")
	}

	season, err := strconv.Atoi(seasonParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	league, err := strconv.Atoi(leagueParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	info := core.Info{
		League: uint(league),
		Season: uint(season),
	}

	standings, err := h.service.GetStandings(info)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(standings)
}

func (h standingsHandler) UpdateStandings(c *fiber.Ctx) error {
	seasonParam := c.Query("season")
	leagueParam := c.Query("league")
	season, err := strconv.Atoi(seasonParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	league, err := strconv.Atoi(leagueParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	info := core.Info{
		League: uint(league),
		Season: uint(season),
	}
	err = h.service.UpdateStandings(info)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"msg": "ok",
	})
}
