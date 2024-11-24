package handler

import (
	"backend/internal/domain/request"
	"backend/internal/domain/response"
	"backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) CreateUserWithTenant(c *fiber.Ctx) error {
	var req request.CreateUserWithTenantRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	tenant, user, err := h.authService.CreateUserWithTenant(c.Context(), &req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"tenant": response.NewTenantResponse(*tenant),
		"user":   response.NewUserResponse(*user),
	})
}
