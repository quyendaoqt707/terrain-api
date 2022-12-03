package controller

import (
	"TerraInnAPI/database"
	"TerraInnAPI/model"
	"TerraInnAPI/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/ulule/deepcopier"
)

func Login(c *fiber.Ctx) error {

	type LoginInput struct {
		Email    string `json:"email" example:"csvadmin"`
		Password string `json:"password" example:"csv12345"`
	}

	//Parse input
	input := new(LoginInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	// check if a user exists
	user := new(model.User)
	db := database.DB

	if res := db.Where("email = ?", input.Email).First(&user); res.RowsAffected <= 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "username_password_incorrect"}) //401
	}
	// Check password
	if err := model.CheckPasswordHash(user.Password, input.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "username_password_incorrect"})
	}

	// Return response
	return c.JSON(fiber.Map{
		"message": "success",
		"token":   utils.GenBase64Token(input.Email),
	})
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

//For profile page
func GetUser(c *fiber.Ctx) error {
	user := new(model.User)
	queryResult := database.DB.Model(model.User{}).Where("email = ?", c.Locals("email").(string)).First(&user)
	if queryResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
	}

	type ReturnStruct struct {
		Id          int    `json:"id"` //row_id
		Email       string `json:"email"`
		Phone       string `json:"phone"`
		FullName    string `json:"full_name"`
		DateOfBirth string `json:"date_of_birth"`
		CidNumber   bool   `json:"cid_number"`
	}
	returnUser := new(ReturnStruct)
	deepcopier.Copy(user).To(returnUser)

	return c.JSON(returnUser)
}

func InsertUser(c *fiber.Ctx) error {

	var err error
	db := database.DB
	user := new(model.User)
	type paramRequest struct {
		Email    string
		Password string
	}

	// param := new(paramRequest)
	// temp, _ := regexp.Compile("[^a-zA-Z0-9]+")

	// param.Username = temp.ReplaceAllString(c.FormValue("userName"), "")
	// param.Password = c.FormValue("password")
	input := new(paramRequest)
	err = c.BodyParser(input)
	if err != nil || len(input.Email) == 0 || len(input.Password) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	// Check exists
	if res := db.Where("email", input.Email).Find(&user); res.RowsAffected > 0 {
		return c.Status(501).JSON(fiber.Map{
			"message": "email_already_exists",
		})
	}

	/**
	*	User
	* ------------------------
	 */
	user.Email = input.Email
	user.Password = model.HashPassword(input.Password)

	// Hash Password and Insert User DB
	// user.UserPassword = model.HashPassword(param.Password)
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})

	}

	// Return response
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func ChangePasswordUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	type paramRequest struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}

	param := new(paramRequest)
	param.CurrentPassword = c.FormValue("current_password")
	param.NewPassword = c.FormValue("new_password")

	//Check password lens
	if len(param.CurrentPassword) < 1 || len(param.NewPassword) < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	//Get user info
	if res := db.Where("email = ?", c.Locals("email").(string)).First(&user); res.RowsAffected <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	// Check password
	if err := model.CheckPasswordHash(user.Password, param.CurrentPassword); err != nil {
		return c.Status(501).JSON(fiber.Map{"message": "current_password_incorrect"})
	}

	// Hash new password and Update new password
	user.Password = model.HashPassword(param.NewPassword)
	db.Save(&user)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// var validate *validator.Validate

func UpdateProfile(c *fiber.Ctx) error {
	db := database.DB

	type UpdateParam struct {
		// Id          int    `json:"id"` //row_ids
		Email       string `json:"email" validate:"required"`
		Phone       string `json:"phone" validate:"required"`
		FullName    string `json:"full_name" validate:"required"`
		DateOfBirth string `json:"date_of_birth" validate:"required"`
		CidNumber   bool   `json:"cid_number" validate:"required"`
	}

	param := new(UpdateParam)

	err := c.BodyParser(param)
	// if err != nil || validate.Struct(param) != nil {
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "bad_request"})
	}

	email := c.Locals("email").(string)

	// err := db.Exec(`UPDATE "tbl_user_info" SET language= ?  WHERE user_name = ?`, language, username)
	err = db.Model(model.UserInfo{}).Where("email = ?", email).Updates(model.User{Email: param.Email, Phone: param.Phone, FullName: param.FullName, DateOfBirth: param.DateOfBirth, CidNumber: param.CidNumber}).Error
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status_code": STATUS_CODE_SUCCESS, "message": "success"})

	}

	// Return response
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
}
