package model

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type EmployeeApprovedStruct struct {
	EmployeeIdArray   string `json:"employee_id_array"`
	EmployeeNameArray string `json:"employee_name_array"`
}
