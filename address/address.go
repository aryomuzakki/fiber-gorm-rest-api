package address

import (
	"errors"
	"fiber-gorm-rest-api/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Address struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Coordinate  string `json:"coordinate"`
}

func GetAll(c *fiber.Ctx) error {
	var addr []Address

	result := database.DB.Find(&addr)

	if result.Error != nil {
		return c.Status(500).JSON(`{"error": "` + result.Error.Error() + `"`)
	}

	return c.Status(200).JSON(addr)
}

func Get(c *fiber.Ctx) error {
	var addr Address

	id, er := strconv.Atoi(c.Params("id"))
	if er != nil {
		return c.Status(400).SendString("id parameter must be an Integer type")
	}

	result := database.DB.First(&addr, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).SendString(result.Error.Error())
		}

		return c.Status(500).JSON(`{"error": "` + result.Error.Error() + `"`)
	}

	return c.Status(200).JSON(addr)
}

func Create(c *fiber.Ctx) error {

	var addr Address

	if err := c.BodyParser(&addr); err != nil {
		return c.Status(400).JSON(`{"error": "` + err.Error() + `"`)
	}

	result := database.DB.Create(&addr)

	if result.Error != nil {
		return c.Status(500).JSON(`{"error": "` + result.Error.Error() + `"`)
	}

	return c.Status(201).JSON(addr)
}

func Update(c *fiber.Ctx) error {

	addr := new(Address)

	if err := c.BodyParser(addr); err != nil {
		return c.Status(400).JSON(`{"error": "` + err.Error() + `"`)
	}

	if addr.ID == 0 {
		return c.Status(400).SendString("ID has not been specified in the request body")
	}
	result := database.DB.Save(&addr)

	if result.Error != nil {
		return c.Status(500).JSON(`{"error": "` + result.Error.Error() + `"`)
	}

	return c.Status(201).JSON(addr)
}

func Delete(c *fiber.Ctx) error {
	var addr Address

	id, er := strconv.Atoi(c.Params("id"))
	if er != nil {
		return c.Status(400).SendString("id parameter must be an Integer type")
	}

	result := database.DB.First(&addr, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).SendString(result.Error.Error())
		}

		return c.Status(500).JSON(`{"error": "` + result.Error.Error() + `"`)
	}

	deleteResult := database.DB.Delete(&addr)

	if deleteResult.Error != nil {
		return c.Status(500).JSON(`{"error": "` + deleteResult.Error.Error() + `"`)
	}

	return c.Status(200).JSON(addr)
}
