package swagger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/redbirdztc/wego/docs"
)

func SwagFunc(ctx *fiber.Ctx) error {
	return swagger.HandlerDefault(ctx)
}
