package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swagger handler
	"github.com/mattiadevivo/crm/controllers"
	_ "github.com/mattiadevivo/crm/docs/crm"
)

func Setup(app *fiber.App) {
	customersRouter := app.Group("/customers")
	customersRouter.Get("/", controllers.GetCustomers)         // get all customers
	customersRouter.Get("/:id", controllers.GetCustomer)       // get single customer
	customersRouter.Post("/", controllers.AddCustomer)         // insert customer
	customersRouter.Put("/:id", controllers.UpdateCustomer)    // modify customer
	customersRouter.Delete("/:id", controllers.DeleteCustomer) // modify customer
	// Others
	app.Static("/", "./static")
	app.Get("/swagger/*", swagger.HandlerDefault)
}
