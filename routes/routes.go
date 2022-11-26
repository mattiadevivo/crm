package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swagger handler
	_ "github.com/mattiadevivo/crm/docs/crm"
)

// HealthCheck godoc
// @Summary Show the status of server. Bananaa
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func dummyController(c *fiber.Ctx) error {
	return nil
}

func Setup(app *fiber.App) {
	app.Get("/*", swagger.HandlerDefault)
	app.Get("/customers/:id", dummyController)    // get single customer
	app.Get("/customers", dummyController)        // get all customers
	app.Post("/customers", dummyController)       // insert customer
	app.Put("/customers/:id", dummyController)    // modify customer
	app.Delete("/customers/:id", dummyController) // modify customer
}
