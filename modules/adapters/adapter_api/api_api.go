package adapterapi

import (
	coreapi "false_api/modules/core/core_api"

	"github.com/gofiber/fiber/v2"
)

type apiHandler struct {
	service coreapi.ApiService
}

func NewApiHandler(service coreapi.ApiService) *apiHandler {
	return &apiHandler{service: service}
}

// func (h apiHandler) CreateTables(c *fiber.Ctx) error {
// 	info := coreapi.Info{}
// 	if err := c.BodyParser(&info); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}
// 	err := h.service.CreateTables(info)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}
// 	return c.JSON(fiber.Map{
// 		"msg": "create ok",
// 	})
// }

func (h apiHandler) CreaatePlayer(c *fiber.Ctx) error {
	info := coreapi.Info{}
	if err := c.BodyParser(&info); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err := h.service.CreatePlayers(info)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"msg": "create ok",
	})
}

func (h apiHandler) CreateMatch(c *fiber.Ctx) error {
	info := coreapi.Info{}
	if err := c.BodyParser(&info); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err := h.service.CreateMatch(info)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"msg": "create ok",
	})
}
