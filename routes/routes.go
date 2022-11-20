package routes

import (
	"github.com/gofiber/fiber/v2"
)

func dummyController(c *fiber.Ctx) error {
	return nil
}

func Setup(app *fiber.App) {
	app.Post("/api/register", dummyController)
}
