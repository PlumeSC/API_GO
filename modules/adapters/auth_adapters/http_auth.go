package authadapter

import (
	authcore "false_api/modules/core/auth_core"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	service authcore.AuthService
}

func NewAuthHandler(service authcore.AuthService) *authHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	user := authcore.Register{}
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

func (h *authHandler) Login(c *fiber.Ctx) error {
	user := authcore.Login{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	token, userInfo, err := h.service.Login(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(fiber.Map{
		"token":     token,
		"user_info": userInfo,
	})
}
