package httpservice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redbirdztc/wego/internal/httpservice/routes/example"
	"github.com/redbirdztc/wego/internal/httpservice/swagger"
)

type Service struct {
	app *fiber.App
}

func New() *Service {
	app := fiber.New(fiber.Config{DisableStartupMessage: false})

	route(app)
	return &Service{
		app: app,
	}
}

func (s *Service) Start(addr string) error {
	return s.app.Listen(addr)
}

func (s *Service) App() *fiber.App {
	return s.app
}

func route(app *fiber.App) {
	app.Get("/", health)
	app.Get("/swagger/*", swagger.SwagFunc)
	apiV1Group := app.Group("api/v1")
	apiV1Group.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
		})
	})

	group1 := apiV1Group.Group("/router")

	example.Route(group1)
}

// health godoc
// @Summary Health check
// @Description Get service health status
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "healthy",
	})
}
