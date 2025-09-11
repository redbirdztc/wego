package example

import (
	"github.com/gofiber/fiber/v2"
)

func Route(router fiber.Router) {
	router.Get("/health", health)
}

// health godoc
// @Summary Health check
// @Description Get service health status
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/group1/health [get]
func health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}
