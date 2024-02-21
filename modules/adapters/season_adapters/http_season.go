package seasonadapters

import (
	core "false_api/modules/core"
	seasoncore "false_api/modules/core/season_core"

	"github.com/gofiber/fiber/v2"
)

type seasonHandler struct {
	service seasoncore.SeasonService
}

func NewSeasonHandler(service seasoncore.SeasonService) *seasonHandler {
	return &seasonHandler{service: service}
}

func (h *seasonHandler) CreateStandings(c *fiber.Ctx) error {
	info := core.Info{}
	if err := c.BodyParser(&info); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err := h.service.CreateStandings(info)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"msg": "ok",
	})
}
func (h *seasonHandler) CreatePlayers(c *fiber.Ctx) error {
	info := core.Info{}
	if err := c.BodyParser(&info); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err := h.service.CreatePlayers(info)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"msg": "ok",
	})
}
func (h *seasonHandler) CreateMatch(c *fiber.Ctx) error {
	info := core.Info{}
	if err := c.BodyParser(&info); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err := h.service.CreateMatch(info)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"msg": "ok",
	})
}
