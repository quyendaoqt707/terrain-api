package controller

import (
	"github.com/gofiber/fiber/v2"
)

const STATUS_CODE_SUCCESS = 0
const STATUS_CODE_FAILURE = 1

func Wellcome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Welcome to Terra API :)))"})
}
