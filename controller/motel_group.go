package controller

import (
	"TerraInnAPI/database"
	"TerraInnAPI/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetModelGroup(c *fiber.Ctx) error {
	id := c.Query("id")
	if len(id) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}
	motel := new(model.MotelGroup)
	rs := database.DB.First(&motel, id) //with primary key
	if rs.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	if rs.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not_found"})

	}

	return c.Status(fiber.StatusOK).JSON(motel)

}

func CreateMotelGroup(c *fiber.Ctx) error {
	var motelGroup model.MotelGroup
	if c.BodyParser(&motelGroup) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	fmt.Printf("%+v", motelGroup)
	motelGroup.OwnerId = c.Locals("email").(string)

	rs := database.DB.Create(&motelGroup) //must be pass an address, otherwise -> panic: reflect.flag.mustBeAssignableSlow(0xc0000b2d00?)
	if rs.Error != nil || rs.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})

}

func UpdateMotelGroup(c *fiber.Ctx) error {
	var motelGroup model.MotelGroup

	if c.BodyParser(&motelGroup) != nil || motelGroup.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}
	motelGroup.OwnerId = c.Locals("email").(string)
	rs := database.DB.Save(motelGroup)
	if rs.Error != nil || rs.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func DelMotelGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	queryResult := database.DB.Delete(&model.MotelGroup{}, id)

	if queryResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	if queryResult.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not_found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})

}
