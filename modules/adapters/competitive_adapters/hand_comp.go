package competitiveadapters

import (
	competitivecore "false_api/modules/core/competitive_core"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type compHandler struct {
	service competitivecore.CompService
}

func NewCompHandler(service competitivecore.CompService) *compHandler {
	return &compHandler{service: service}
}

func (h *compHandler) GetMatchDay(c *fiber.Ctx) error {
	err := h.service.Comp()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.JSON("ok")
}

func (h *compHandler) Live() {
	err := h.service.Comp()
	if err != nil {
		fmt.Println(err)
	}
}
