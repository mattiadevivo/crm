package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mattiadevivo/crm/database"
	"github.com/mattiadevivo/crm/models"
	"github.com/rs/zerolog/log"
)

func GetCustomers(c *fiber.Ctx) error {
	log.Debug().Msg("GET /customers")
	var customers []models.Customer
	result := database.DB.Find(&customers)
	if result.Error != nil {
		log.Error().Err(result.Error)
		return c.Status(404).JSON(&fiber.Map{
			"error": "No customers found!",
		})
	}
	return c.JSON(customers)
}

func GetCustomer(c *fiber.Ctx) error {
	log.Debug().Msgf("GET /customers/%s", c.Params("id"))
	var customer models.Customer
	result := database.DB.First(&customer, c.Params("id"))
	if result.Error != nil {
		log.Error().Err(result.Error)
		return c.Status(404).JSON(&fiber.Map{
			"error": fmt.Sprintf("No customer with id %s found!", c.Params("id")),
		})
	}
	return c.JSON(customer)
}

func AddCustomer(c *fiber.Ctx) error {
	var customer models.Customer
	c.Accepts("application/json")
	c.AcceptsEncodings("encoding/json")
	if err := c.BodyParser(&customer); err != nil {
		log.Error().Err(err)
		c.Status(400).JSON(&fiber.Map{
			"error": "Invalid customer received",
		})
	}
	log.Debug().Interface("customer", customer).Msg("POST /customers")
	// Insert new customer inside DB
	if result := database.DB.Create(&customer); result.Error != nil {
		log.Error().Err(result.Error)
		c.Status(500).JSON(&fiber.Map{
			"error": fmt.Sprintf("Error while adding customer to db: %s", result.Error),
		})
	}
	return c.JSON(customer)
}

func UpdateCustomer(c *fiber.Ctx) error {
	var customer models.Customer
	var updatedCustomer models.Customer
	c.Accepts("application/json")
	c.AcceptsEncodings("encoding/json")
	log.Debug().Msgf("PUT /customers/%s", c.Params("id"))
	result := database.DB.First(&customer, c.Params("id"))
	if result.Error != nil {
		log.Error().Err(result.Error)
		return c.Status(404).JSON(&fiber.Map{
			"error": fmt.Sprintf("No customer with id %s found!", c.Params("id")),
		})
	}
	if err := c.BodyParser(&updatedCustomer); err != nil {
		log.Error().Err(err)
		c.Status(400).JSON(&fiber.Map{
			"error": "Invalid customer received",
		})
	}
	// Update Customer with the new one
	id, _ := strconv.Atoi(c.Params("id"))
	updatedCustomer.Id = id
	if result := database.DB.Save(&updatedCustomer); result.Error != nil {
		log.Error().Err(result.Error)
		c.Status(500).JSON(&fiber.Map{
			"error": fmt.Sprintf("Error while upadting customer to db: %s", result.Error),
		})
	}
	return c.JSON(updatedCustomer)
}

func DeleteCustomer(c *fiber.Ctx) error {
	log.Debug().Msgf("DELETE /customers/%s", c.Params("id"))
	var customer models.Customer
	result := database.DB.First(&customer, c.Params("id"))
	if result.Error != nil {
		log.Error().Err(result.Error)
		return c.Status(404).JSON(&fiber.Map{
			"error": fmt.Sprintf("No customer with id %s found!", c.Params("id")),
		})
	}
	// Delete customer from DB
	if result := database.DB.Delete(&customer); result.Error != nil {
		log.Error().Err(result.Error)
		c.Status(500).JSON(&fiber.Map{
			"error": fmt.Sprintf("Error while deleting customer from db: %s", result.Error),
		})
	}
	return c.JSON(&fiber.Map{
		"message": fmt.Sprintf("Correctly deleted customer with id %s", c.Params("id")),
	})
}
