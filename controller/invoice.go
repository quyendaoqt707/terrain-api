package controller

import (
	"TerraInnAPI/database"
	"TerraInnAPI/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetInvoice(c *fiber.Ctx) error {
	//Get detail
	//Get list
	id := c.Query("id")
	if id != "" {
		return getInvoiceDetail(c, id)
	}

	groupId := c.Query("group-id")
	month := c.Query("month")
	if groupId != "" && month != "" {
		return getListInvoiceByMonth(c, groupId, month)
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})

}

func getInvoiceDetail(c *fiber.Ctx, id string) error {
	var returnObject struct {
		model.Invoice
		RoomName string `json:"room_name"`
	}
	// invoice := new(model.Invoice)

	// rs := database.DB.First(&invoice, id) //with primary key
	rs := database.DB.Table("invoices").Select("invoices.*, motel.name AS room_name").
		Joins("LEFT JOIN motel ON motel.id = invoices.motel_id").
		Where("invoices.id = ?", id).Scan(&returnObject)
	if rs.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	if rs.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not_found"})

	}
	// return c.Status(fiber.StatusOK).JSON(invoice)
	return c.Status(fiber.StatusOK).JSON(returnObject)

}

func getListInvoiceByMonth(c *fiber.Ctx, groupId string, month string) error {
	var returnStruct []struct {
		Id        int    `json:"id"`
		RoomName  string `json:"room_name"`
		ElecUsed  int    `json:"elec_used"`
		WaterUsed int    `json:"water_used"`
	}

	sql := fmt.Sprintf(`
	SELECT invoices.id, elec_index_after - elec_index_before AS  elec_used,
	water_index_after - water_index_before AS water_used,
	motel.name AS room_name
	FROM invoices LEFT JOIN motel ON motel.id = invoices.motel_id
	WHERE motel.group_id =1 AND invoice_date ='%s'`, month) //2022-12
	rs := database.DB.Raw(sql).Scan(&returnStruct)

	if rs.Error != nil {
		fmt.Println(rs.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})

	}

	return c.JSON(returnStruct)
}

func CreateInvoice(c *fiber.Ctx) error {
	var invoice model.Invoice
	if c.BodyParser(&invoice) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	// fmt.Printf("%+v", invoice)
	// motelGroup.OwnerId = c.Locals("email").(string)

	rs := database.DB.Create(&invoice) //must be pass an address, otherwise -> panic: reflect.flag.mustBeAssignableSlow(0xc0000b2d00?)
	if rs.Error != nil || rs.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func UpdateInvoice(c *fiber.Ctx) error {
	var invoice model.Invoice

	if c.BodyParser(&invoice) != nil || invoice.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}
	// invoice.OwnerId = c.Locals("email").(string)
	rs := database.DB.Save(invoice)
	if rs.Error != nil || rs.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func DelInvoice(c *fiber.Ctx) error {
	id := c.Params("id")
	queryResult := database.DB.Delete(&model.Invoice{}, id)

	if queryResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	if queryResult.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not_found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})

}
