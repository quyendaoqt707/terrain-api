package controller

import (
	"TerraInnAPI/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetPositionList(c *fiber.Ctx) error {
	type SalaryStruct struct {
		Id    int    `json:"id"`
		Value string `json:"value"`
	}
	type PositionStructWithRankList struct {
		PositionId    int            `json:"position_id"`
		PositionName  string         `json:"position_name"`
		SlaryRankList []SalaryStruct `json:"salary_rank_list"`
	}

	type PositionStruct struct {
		PositionId   int    `json:"position_id"`
		PositionName string `json:"position_name"`
	}

	var positionList []PositionStruct
	var positionStructWithRankList []PositionStructWithRankList

	var scanStruct []struct {
		PositionId   int    `json:"position_id"`
		PositionName string `json:"position_name"`
		SalaryRankId int    `json:"salary_rank_id"`
		ContentVn    string `json:"content_vn"`
	}

	db := database.DB
	var queryResult *gorm.DB
	mode := false
	if c.Query("with-salary-rank-list") == "true" {
		mode = true
		sql := `
		SELECT distinct tbl_position.position_id, position_name, h_salary_rank.id AS salary_rank_id , h_salary_rank.content_vn as content_vn
		FROM tbl_position LEFT JOIN h_review_role ON tbl_position.position_id = h_review_role.position_id
		LEFT JOIN h_salary_rank ON salary_rank_id = h_salary_rank.id
		WHERE delete_fg = 0
		ORDER BY tbl_position.position_id`
		queryResult = db.Raw(sql).Scan(&scanStruct)

		//Procesing data

		// var a []SalaryStruct
		slaryRankMap := make(map[int][]SalaryStruct)
		positionNameMap := make(map[int]string)
		for _, item := range scanStruct {
			// _, isExist := slaryRankMap[item.PositionId]
			// if isExist {
			if item.SalaryRankId != 0 {
				slaryRankMap[item.PositionId] = append(slaryRankMap[item.PositionId], SalaryStruct{Id: item.SalaryRankId, Value: item.ContentVn})
			}

			// } else {
			// 	slaryRankMap[item.PositionId] =
			// }
			positionNameMap[item.PositionId] = item.PositionName

		}

		for key, value := range slaryRankMap {
			positionStructWithRankList = append(positionStructWithRankList, PositionStructWithRankList{PositionId: key, PositionName: positionNameMap[key], SlaryRankList: value})
		}

	} else {
		queryResult = db.Table("tbl_position").Select("position_id", "position_name").Where("delete_fg = 0 ").Scan(&positionList)
	}

	if queryResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
	}

	if mode {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status_code": STATUS_CODE_SUCCESS, "message": "success", "list": positionStructWithRankList})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status_code": STATUS_CODE_SUCCESS, "message": "success", "list": positionList})
}

func GetSalaryRankList(c *fiber.Ctx) error {
	type ReturnStruct struct {
		SalaryRankVn []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"lst_quota_vn"`
		SalaryRankEn []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"lst_quota_en"`
		SalaryRankJp []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"lst_quota_jp"`
	}
	var salaryRankList ReturnStruct
	db := database.DB
	var salaryRankListModel []struct {
		Id        int    `json:"id"`
		ContentVn string `json:"content_vn"`
		ContentEn string `json:"content_en"`
		ContentJp string `json:"content_Jp"`
	}
	queryResult := db.Table("h_salary_rank").Select("id", "content_vn", "content_en", "content_jp").Where("del_flg = 0 ").Scan(&salaryRankListModel)

	if queryResult.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
	}

	for _, rank := range salaryRankListModel {
		salaryRankList.SalaryRankEn = append(salaryRankList.SalaryRankEn, struct {
			Id   int    "json:\"id\""
			Name string "json:\"name\""
		}{Id: rank.Id, Name: rank.ContentEn})
		salaryRankList.SalaryRankVn = append(salaryRankList.SalaryRankVn, struct {
			Id   int    "json:\"id\""
			Name string "json:\"name\""
		}{Id: rank.Id, Name: rank.ContentVn})
		salaryRankList.SalaryRankJp = append(salaryRankList.SalaryRankJp, struct {
			Id   int    "json:\"id\""
			Name string "json:\"name\""
		}{Id: rank.Id, Name: rank.ContentJp})
	}

	return c.Status(fiber.StatusOK).JSON(salaryRankList)
}

func PrepareDataForFlowSetting(c *fiber.Ctx) error {
	type FormType struct {
		FormTypeName string `json:"name"`
		FormTypeId   int    `json:"id"`
	}
	type EmployeeList struct {
		EmployeeName string `json:"name"`
		EmployeeId   string `json:"id"`
	}
	type DeptList struct {
		DepartmentId   int    `json:"id"`
		DepartmentName string `json:"name"`
	}

	type TeamList struct {
		TeamId   int    `json:"id"`
		TeamName string `json:"name"`
	}

	type GroupList struct {
		GroupId   int    `json:"id"`
		GroupName string `json:"name"`
	}

	type PositionList struct {
		PositionId   int    `json:"id"`
		PositionName string `json:"name"`
	}

	type returnDataType struct {
		FormType     []FormType     `json:"form_type"`
		DeptList     []DeptList     `json:"dept_list"`
		TeamList     []TeamList     `json:"team_list"`
		GroupList    []GroupList    `json:"group_list"`
		PositionList []PositionList `json:"position_list"`
		EmployeeList []EmployeeList `json:"employee_list"`
	}
	returnData := new(returnDataType)

	returnData.FormType = append(returnData.FormType, FormType{FormTypeId: 1, FormTypeName: "Đánh giá"})
	// returnData.FormType = append(returnData.FormType, FormType{FormTypeId: 2, FormTypeName: "FORM TYPE 2"})

	db := database.DB
	var groupList []GroupList
	queryResult := db.Table("tbl_group").Select("group_id", "group_name").Where("delete_fg = 0 ").Scan(&groupList)
	if queryResult.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
	}
	returnData.GroupList = groupList

	var teamList []TeamList
	queryResult = db.Table("tbl_team").Select("team_id", "team_name").Where("delete_fg = 0 ").Scan(&teamList)
	if queryResult.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
	}
	returnData.TeamList = teamList

	var deptList []DeptList
	queryResult = db.Table("tbl_department").Select("department_id", "department_name").Where("delete_fg = 0 ").Scan(&deptList)
	if queryResult.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
	}
	returnData.DeptList = deptList

	var positionList []PositionList
	queryResult = db.Table("tbl_position").Select("position_id", "position_name").Where("delete_fg = 0 ").Scan(&positionList)
	if queryResult.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
	}
	returnData.PositionList = positionList

	var employeeList []EmployeeList
	queryResult = db.Table("tbl_employeeprofile").Select("employee_id", "employee_name").Where("status_id = 1 ").Scan(&employeeList)
	if queryResult.Error != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
	}
	returnData.EmployeeList = employeeList

	return c.Status(fiber.StatusOK).JSON(returnData)
}

func GetEmployeeList(c *fiber.Ctx) error {
	type EmployeeList struct {
		EmployeeName string `json:"name"`
		EmployeeId   string `json:"id"`
	}

	var employeeList []EmployeeList
	queryResult := database.DB.Table("tbl_employeeprofile").Select("employee_id", "employee_name").Where("status_id = 1 ").Scan(&employeeList) //--> incorrect
	// $this->db->where("(retired.retired_status is null or retired.retired_status = 0 or retired.retired_date > '$now')"); --> correct (php)
	// queryResult := database.DB.Table()
	if queryResult.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status_code": STATUS_CODE_FAILURE, "message": "system_error"})
	}

	return c.Status(fiber.StatusOK).JSON(employeeList)
}

// func WriteWebLog(c *fiber.Ctx) error {
// 	log := new(model.WebLog)
// 	input := c.BodyParser(log)
// 	if input == nil {
// 		log.UserName = c.Locals("username").(string)
// 		database.DB.Create(log)
// 		c.Status(fiber.StatusOK)
// 		return nil
// 	}
// 	c.Status(fiber.StatusInternalServerError).SendString(input.Error())
// 	return nil
// }
