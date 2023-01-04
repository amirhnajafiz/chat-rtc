package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Welcome will render the welcome page template.
func (h *Handler) Welcome(c *fiber.Ctx) error {
	return c.Render("welcome", nil, "layouts/main")
}
