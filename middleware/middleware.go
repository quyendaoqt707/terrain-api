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
	email, err := utils.ExtractBase64Token(token)
	if err != nil || email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status_code": 1, "message": "token_invalid"})
	}

	// Check user exists
	if res := db.Where("email = ?", email).First(&user); res.RowsAffected <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status_code": 1, "message": "token_invalid"})

	}
	// Set data log
	c.Locals("email", email)
	return c.Next()
}
