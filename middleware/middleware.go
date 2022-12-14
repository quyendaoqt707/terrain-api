package middleware

import (
	"TerraInnAPI/database"
	"TerraInnAPI/model"
	"TerraInnAPI/utils"

	"github.com/gofiber/fiber/v2"
)

// Authen
func AppAuthen(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	// Check token valid
	token := c.Get("token")
	phone, err := utils.ExtractBase64Token(token)
	if err != nil || phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status_code": 1, "message": "token_invalid"})
	}

	// Check user exists
	if res := db.Where("phone = ?", phone).First(&user); res.RowsAffected <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status_code": 1, "message": "token_invalid"})

	}
	// Set data log
	c.Locals("phone", phone)
	return c.Next()
}
