package controller

import (
	"TerraInnAPI/database"
	"TerraInnAPI/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func getMotelListForGuest(c *fiber.Ctx) error {
	var motelList []struct {
		GroupName string `json:"group_name"`
		model.Motel
	}
	// var motelList []model.Motel
	// rs := database.DB.Find(&motelList) //with primary key

	rs := database.DB.Model(model.Motel{}).Select("motel.* , group_name").Joins("JOIN motel_group ON motel_group.id = motel.group_id").Scan(&motelList)

	if rs.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}
	return c.Status(fiber.StatusOK).JSON(motelList)
}

func getMotelList(c *fiber.Ctx) error {
	var motelList []model.Motel

	groupId := c.Query("group-id")
	if len(groupId) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	rs := database.DB.Find(&motelList, " group_id = "+groupId) //with primary key
	if rs.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	// if rs.RowsAffected == 0 {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not_found"})
	// }

	return c.Status(fiber.StatusOK).JSON(motelList)
}

func GetModelDetail(c *fiber.Ctx) error {
	forGuest := c.Query("guest")
	if forGuest == "true" {
		return getMotelListForGuest(c)
	}

	groupId := c.Query("group-id")
	if len(groupId) != 0 {
		return getMotelList(c)
	}

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
