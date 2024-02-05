package adaptersauth

import (
	coreauth "false_api/modules/core/core_auth"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	service coreauth.AuthService
}

func NewAuthHandler(service coreauth.AuthService) *authHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	user := coreauth.Register{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	token, userInfo, err := h.service.Register(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"token":     token,
		"user_info": userInfo,
	})
}
