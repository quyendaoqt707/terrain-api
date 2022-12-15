package controller

import (
	"TerraInnAPI/database"
	"TerraInnAPI/model"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadImg(c *fiber.Ctx) error {

	// fileHeader, _ := c.FormFile("files")
	// c.SaveFile(fileHeader, "/assets")

	formKeys, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	var imgIds []string
	i := 0
	for formFieldName, fileHeaders := range formKeys.File {
		fmt.Println(formFieldName)

		for _, fileHeader := range fileHeaders {
			// fmt.Println(fileHeader)
			// fileHeader.Header.Get("Content-Type")

			if len(strings.Split(fileHeader.Filename, ".")) < 2 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "filename_invalid"})

			}

			fileExtension := strings.Split(fileHeader.Filename, ".")[len(strings.Split(fileHeader.Filename, "."))-1]
			fileName := fmt.Sprintf("img_%d%d.%s", time.Now().Unix(), i, fileExtension)
			fmt.Println("fileName: ", fileName, " fileSize=", fileHeader.Size, " byte")
			err := c.SaveFile(fileHeader, "./assets/"+fileName)
			if err != nil {
				fmt.Println(err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "system_error"})

			}

			imgIds = append(imgIds, fileName)
			i = i + 1
		}
		i = i + 1
	}

	// files := c.FormValue("files")
	// fmt.Println(files)

	return c.JSON(fiber.Map{"img_id": imgIds})
}

func AcceptJoinRq(c *fiber.Ctx) error {
	db := database.DB

	type UpdateParam struct {
		RequestId int `json:"request_id"`
		MotelId   int `json:"motel_id"`
		// Creator   string `json:"creator"`
		Action int `json:"action"` //1: Accept/ 2: Deny 3: Mark as done (sửa chữa)
	}

	param := new(UpdateParam)

	err := c.BodyParser(param)
	// if err != nil || validate.Struct(param) != nil {
	if err != nil || param.Action == 0 || param.RequestId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "param_invalid"})
	}

	if (param.Action == 1 || param.Action == 2) && param.MotelId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "param_invalid"})

	}
	// phone := c.Locals("phone").(string)

	request := new(model.Request)
	rs := db.Find(&request, param.RequestId)
	if rs.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "request_id_invalid"})
	}

	// if request.RequestType ==1 { //Trả hoặc thuê phòng
	// }

	if param.Action == 1 {
		// err := db.Exec(`UPDATE "tbl_user_info" SET language= ?  WHERE user_name = ?`, language, username)
		rs := db.Model(model.User{}).Where("phone = ?", request.Creator).Updates(model.User{MotelId: param.MotelId}) //Where phải trước Updates
		if rs.Error == nil && rs.RowsAffected > 0 {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status_code": STATUS_CODE_SUCCESS, "message": "success"})
		}
	}

	if param.Action == 2 || param.Action == 3 {
		request.Status = 2 //Đã hoàn thành
		rs := db.Save(request)
		if rs.Error == nil && rs.RowsAffected > 0 {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status_code": STATUS_CODE_SUCCESS, "message": "success"})
		}
	}

	// Return response
	// return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "bad_request"})

}
