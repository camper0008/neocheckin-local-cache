package response_models

import im "neocheckin_cache/wrapper/models/imported_models"

type GetEmployees struct {
	Data []im.Employee `json:"data"`
}

type GetEmployeesSync struct {
	Data []im.EmployeeSync `json:"data"`
}
