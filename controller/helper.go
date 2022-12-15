package controller

import (
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
