package controller

import (
	"TerraInnAPI/database"
	"TerraInnAPI/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetModelDetail(c *fiber.Ctx) error {
	id := c.Query("id")
	if len(id) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}
	motel := new(model.Motel)
	rs := database.DB.First(&motel, id) //with primary key
	if rs.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	if rs.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not_found"})

	}

	return c.Status(fiber.StatusOK).JSON(motel)

}

func CreateMotel(c *fiber.Ctx) error {
	var motel model.Motel
	if c.BodyParser(&motel) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	fmt.Printf("%+v", motel)

	rs := database.DB.Create(&motel) //must be pass an address, otherwise -> panic: reflect.flag.mustBeAssignableSlow(0xc0000b2d00?)
	if rs.Error != nil || rs.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})

}

func UpdateMotel(c *fiber.Ctx) error {
	var motel model.Motel

	if c.BodyParser(&motel) != nil || motel.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	rs := database.DB.Save(motel)
	if rs.Error != nil || rs.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func DelMotel(c *fiber.Ctx) error {
	id := c.Params("id")
	queryResult := database.DB.Delete(&model.Motel{}, id)

	if queryResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	if queryResult.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not_found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})

}
