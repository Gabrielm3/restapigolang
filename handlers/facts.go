package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gabrielm3/restapigolang/database"
	"github.com/gabrielm3/restapigolang/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
	id := c.Params("id")

	fact := new(models.Fact)
	result := database.DB.Db.First(&fact, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Save(&fact)

	return c.Status(fiber.StatusOK).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {

	id := c.Params("id")

	result := database.DB.Db.Delete(&models.Fact{}, id)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
