package controller

import (
	"TerraInnAPI/database"
	"TerraInnAPI/model"
	"TerraInnAPI/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type userMoreStruct struct {
	model.User
	GroupName string `json:"group_name"`
	RoomName  string `json:"room_name"`
	IsRented  bool   `json:"is_rented"`
}

func Login(c *fiber.Ctx) error {

	type LoginInput struct {
		Phone    string `json:"phone" example:"csvadmin"`
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

	if res := db.Where("phone = ?", input.Phone).First(&user); res.RowsAffected <= 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "phone_or_password_incorrect"}) //401
	}
	// Check password
	if err := utils.CheckPasswordHash(user.Password, input.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "phone_or_password_incorrect"})
	}

	// Return response
	return c.JSON(fiber.Map{
		"is_admin": user.IsAdmin,
		"message":  "success",
		"token":    utils.GenBase64Token(input.Phone),
	})
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

//For profile page
func GetUser(c *fiber.Ctx) error {

	if c.Query("motel-id") != "" {
		return GetUserByRoom(c)
	}

	var userMore []userMoreStruct
	queryResult := database.DB.Model(model.User{}).
		Select("user.*, group_name, name as room_name").
		Joins("LEFT JOIN motel ON motel.id = user.motel_id").
		Joins("LEFT JOIN motel_group ON motel.group_id = motel_group.id").
		Where("phone = ?", c.Locals("phone").(string)).
		Scan(&userMore)
	if queryResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
	}

	for idx, item := range userMore {
		if item.MotelId != 0 {
			userMore[idx].IsRented = true
		}
	}

	return c.JSON(userMore)
}

func GetUserByRoom(c *fiber.Ctx) error {
	motelId := c.Query("motel-id")
	// var users []model.User

	type UserByRoom struct {
		Name        string `json:"name"`
		Phone       string `json:"phone"`
		Age         int    `json:"age"`
		DateOfBirth string `json:"date_of_birth"`
		Image       string `json:"image"`
	}

	users := []UserByRoom{} //return empty list instead null
	// queryResult := database.DB.Model(model.User{}).Where("motel_id = ?", motelId).Find(&users)
	queryResult := database.DB.Model(model.User{}).
		Select("full_name, phone, avatar_url as image, date_of_birth").
		Where("motel_id = ?", motelId).Scan(&users)

	if queryResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
	}

	for idx, item := range users {
		year, err := strconv.Atoi(item.DateOfBirth[:4])
		if err == nil {
			users[idx].Age = 2022 - year
		}

	}

	return c.JSON(users)
}

func InsertUser(c *fiber.Ctx) error {

	var err error
	db := database.DB
	user := new(model.User)
	type paramRequest struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}

	// param := new(paramRequest)
	// temp, _ := regexp.Compile("[^a-zA-Z0-9]+")

	// param.Username = temp.ReplaceAllString(c.FormValue("userName"), "")
	// param.Password = c.FormValue("password")
	input := new(paramRequest)
	err = c.BodyParser(input)
	if err != nil || len(input.Phone) == 0 || len(input.Password) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	// Check exists
	if res := db.Where("phone", input.Phone).Find(&user); res.RowsAffected > 0 {
		return c.Status(501).JSON(fiber.Map{
			"message": "phone_already_exists",
		})
	}

	/**
	*	User
	* ------------------------
	 */
	user.Phone = input.Phone
	user.Password = utils.HashPassword(input.Password)
	user.IsAdmin = input.IsAdmin

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

func ChangePassword(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	type paramRequest struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}

	param := new(paramRequest)
	// param.CurrentPassword = c.FormValue("current_password")
	// param.NewPassword = c.FormValue("new_password")

	//Check password lens
	if err := c.BodyParser(&param); err != nil || len(param.CurrentPassword) < 1 || len(param.NewPassword) < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	//Get user info
	if res := db.Where("phone = ?", c.Locals("phone").(string)).First(&user); res.RowsAffected <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad_request"})
	}

	// Check password
	if err := utils.CheckPasswordHash(user.Password, param.CurrentPassword); err != nil {
		return c.Status(501).JSON(fiber.Map{"message": "current_password_incorrect"})
	}

	// Hash new password and Update new password
	user.Password = utils.HashPassword(param.NewPassword)
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
		CidNumber   string `json:"cid_number" validate:"required"`
		MotelId     int    `json:"motel_id"`
		AvatarUrl   string `json:"avatar_url"`
	}

	param := new(UpdateParam)

	err := c.BodyParser(param)
	// if err != nil || validate.Struct(param) != nil {
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "bad_request"})
	}
	phone := c.Locals("phone").(string)

	// err := db.Exec(`UPDATE "tbl_user_info" SET language= ?  WHERE user_name = ?`, language, username)
	rs := db.Model(model.User{}).Where("phone = ?", phone).Updates(model.User{Phone: param.Phone, Email: param.Email, FullName: param.FullName, DateOfBirth: param.DateOfBirth, CidNumber: param.CidNumber}) //Where phải trước Updates
	if rs.Error == nil && rs.RowsAffected > 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status_code": STATUS_CODE_SUCCESS, "message": "success", "new_token": utils.GenBase64Token(param.Phone)})

	}
	// Return response
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
}
