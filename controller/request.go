package controller

import (
	"TerraInnAPI/database"
	"TerraInnAPI/model"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetRequest(c *fiber.Ctx) error {
	//Get detail
	//Get list
	id := c.Query("id")
	if id != "" {
		return getRequestById(c, id)
	}

	roomId := c.Query("room-id")
	// month := c.Query("month")
	if roomId != "" {
		return getRequestByRoom(c, roomId)
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})

}

func getRequestById(c *fiber.Ctx, id string) error {
	type ReturnStruct struct {
		Id          int    `json:"id" gorm:"primaryKey; autoIncreament"`
		Creator     string `json:"creator"` //Creator
		CreatorName string `json:"creator_name"`
		RoomName    string `json:"room_name"`
		GroupName   string `json:"group_name"`
		IsFromAdmin bool   `json:"is_from_admin"`
		MotelId     int    `json:"motel_id"`
		RequestType int    `json:"type"` //Thuee phong//tra phong
		// TypeName    string    `json:"type_name"`
		Status   int       `json:"status"` //0
		Title    string    `json:"title"`
		DueDate  string    `json:"due_date"`
		Content  string    `json:"content"`
		CreateAt time.Time `json:"create_at" gorm:"autoCreateTime"`
	}
	returnStruct := new(ReturnStruct)
	// request := new(model.Request)

	// rs := database.DB.First(&request, id) //with primary key
	sql := fmt.Sprintf(`
	SELECT request.*,user.full_name AS creator_name, motel.name AS room_name, motel_group.group_name FROM request
LEFT JOIN user ON user.phone = request.creator
LEFT JOIN motel ON motel.id = request.motel_id
LEFT JOIN motel_group ON motel_group.id = motel.group_id
WHERE request.id = %s`, id)
	rs := database.DB.Raw(sql).Scan(&returnStruct)
	// deepcopier.Copy(request).To(returnStruct)

	if rs.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	if rs.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not_found"})

	}
	return c.Status(fiber.StatusOK).JSON(returnStruct)
}

func getRequestByRoom(c *fiber.Ctx, roomId string) error {
	var returnStruct []model.Request

	rs := database.DB.Where("motel_id = ? ", roomId).Find(&returnStruct)

	if rs.Error != nil {
		fmt.Println(rs.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})

	}

	return c.JSON(returnStruct)
}

func CreateRequest(c *fiber.Ctx) error {
	var request model.Request
	if c.BodyParser(&request) != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	// fmt.Printf("%+v", invoice)
	// motelGroup.OwnerId = c.Locals("phone").(string)

	rs := database.DB.Create(&request) //must be pass an address, otherwise -> panic: reflect.flag.mustBeAssignableSlow(0xc0000b2d00?)
	if rs.Error != nil || rs.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func UpdateRequest(c *fiber.Ctx) error {
	var request model.Request

	if c.BodyParser(&request) != nil || request.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}
	// invoice.OwnerId = c.Locals("phone").(string)
	rs := database.DB.Save(request)
	if rs.Error != nil || rs.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func DelRequest(c *fiber.Ctx) error {
	id := c.Params("id")
	queryResult := database.DB.Delete(&model.Request{}, id)

	if queryResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})
	}

	if queryResult.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "not_found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})

}
